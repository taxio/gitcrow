package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rakyll/statik/fs"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	_ "github.com/taxio/gitcrow/cmd/gitcrow-cli/statik"
	"golang.org/x/xerrors"
)

var (
	ErrConfigFileAlreadyExists = xerrors.New("config file already exists")
)

type Config struct {
	ServerHost        string
	Username          string
	GitHubAccessToken string
}

type Manager interface {
	Exists() (bool, error)
	GenerateFromTemplate(host, username, token string) error
	Load() (*Config, error)
}

func NewManager(fs afero.Fs) Manager {
	return &managerImpl{fs: fs}
}

type managerImpl struct {
	fs afero.Fs
}

func (c *managerImpl) Exists() (bool, error) {
	ext, err := c.ConfigFileExists()
	if err != nil {
		return false, err
	}
	return ext, nil
}

func (c *managerImpl) ConfigFileExists() (bool, error) {
	configFilePath, err := c.getConfigFilePath()
	if err != nil {
		return false, err
	}
	af := afero.Afero{Fs: c.fs}
	ext, err := af.Exists(configFilePath)
	if err != nil {
		return false, err
	}
	return ext, nil
}

func (c *managerImpl) ConfigDirExists() (bool, error) {
	configBaseDir, err := c.getConfigDirPath()
	if err != nil {
		return false, err
	}
	af := afero.Afero{Fs: c.fs}
	ext, err := af.DirExists(configBaseDir)
	if err != nil {
		return false, err
	}
	return ext, nil
}

func (c *managerImpl) GenerateFromTemplate(host, username, token string) (err error) {
	ext, err := c.Exists()
	if err != nil {
		return err
	}
	if ext {
		return ErrConfigFileAlreadyExists
	}

	cfgData := Config{
		ServerHost:        host,
		Username:          username,
		GitHubAccessToken: token,
	}

	// get template from statik
	statikFs, err := fs.New()
	if err != nil {
		return err
	}
	tplFile, err := statikFs.Open("/config.toml.tmpl")
	if err != nil {
		return err
	}
	defer func() {
		if cErr := tplFile.Close(); err == nil {
			err = cErr
		}
	}()
	b, err := ioutil.ReadAll(tplFile)
	tplStr := string(b)

	// generate config file from template
	// panic when cannot parse template
	tpl := template.Must(template.New("config").Parse(tplStr))
	if err != nil {
		return err
	}
	appConfigFilePath, err := c.getConfigFilePath()
	if err != nil {
		return err
	}
	cfgFile, err := c.fs.Create(appConfigFilePath)
	if err != nil {
		return err
	}
	err = tpl.Execute(cfgFile, cfgData)
	if err != nil {
		return err
	}

	return nil
}

func (c *managerImpl) Load() (*Config, error) {
	appConfigPath, err := c.getConfigFilePath()
	if err != nil {
		return nil, xerrors.Errorf("getAppConfigDirPath: %v", err)
	}

	v := viper.New()
	v.SetFs(c.fs)
	v.SetConfigType("toml")
	v.SetConfigFile(appConfigPath)
	err = v.ReadInConfig()
	if err != nil {
		return nil, xerrors.Errorf("viper.ReadInConfig: %v", err)
	}
	var cfg *Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, xerrors.Errorf("viper.Unmarshal: %v", err)
	}

	return cfg, nil
}

func (c *managerImpl) getConfigFilePath() (string, error) {
	appConfigDirPath, err := c.getAppConfigDirPath()
	if err != nil {
		return "", err
	}
	configFilePath := filepath.Join(appConfigDirPath, "config.toml")
	return configFilePath, nil
}

func (c *managerImpl) getAppConfigDirPath() (string, error) {
	configBaseDir, err := c.getConfigDirPath()
	if err != nil {
		return "", err
	}
	appConfigDirPath := filepath.Join(configBaseDir, "gitcrow")
	return appConfigDirPath, nil
}

func (c *managerImpl) getConfigDirPath() (string, error) {
	configBaseDir := os.Getenv("XDG_CONFIG_HOME")
	if len(configBaseDir) == 0 {
		//log.Println("XDG_CONFIG_HOME not found, use HOME instead.")
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configBaseDir = filepath.Join(homedir, ".config")
	}

	return configBaseDir, nil
}
