package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"golang.org/x/xerrors"
)

var (
	ErrConfigDirNotExists = xerrors.New("config directory not exists")
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
	return &managerImpl{fs: fs, af: afero.Afero{Fs: fs}}
}

type managerImpl struct {
	fs afero.Fs
	af afero.Afero
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
	ext, err := c.af.Exists(configFilePath)
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
	ext, err := c.af.DirExists(configBaseDir)
	if err != nil {
		return false, err
	}
	return ext, nil
}

func (c *managerImpl) GenerateFromTemplate(host, username, token string) error {
	return nil
}

func (c *managerImpl) Load() (*Config, error) {
	return &Config{}, nil
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
		log.Println("XDG_CONFIG_HOME not found, use HOME instead.")
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configBaseDir = filepath.Join(homedir, ".config")
	}

	return configBaseDir, nil
}
