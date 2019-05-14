package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rakyll/statik/fs"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Username          string
	GitHubAccessToken string
}

var (
	ErrAppConfigValidation = errors.New("input your information to config")
	ErrAppConfigNotExists  = errors.New("app config is not exists")
)

const (
	appName            = "gitcrow"
	appConfigName      = "gitcrow.toml"
	configTemplatePath = "/gitcrow.toml.tmpl"
)

func Load() (*Config, error) {
	configFilePath, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	err = existsConfig(configFilePath)
	if err != nil {
		return nil, err
	}

	configHome, _ := filepath.Split(configFilePath)

	// read config
	var cfg *Config
	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigName(appName)
	v.AddConfigPath(configHome)
	err = v.ReadInConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// validate config
	err = validateConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func CreateConfigFile() (string, error) {
	configFilePath, err := getConfigPath()
	if err != nil {
		return "", err
	}
	configFilePath = filepath.Clean(configFilePath)
	ss := strings.Split(configFilePath, string(os.PathSeparator))
	p := "/"
	for _, s := range ss[1 : len(ss)-1] {
		p = filepath.Join(p, s)
		_, err := os.Stat(p)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.Mkdir(p, 0755)
				if err != nil {
					return "", errors.WithStack(err)
				}
				fmt.Printf("mkdir %s\n", p)
			} else {
				return "", errors.WithStack(err)
			}
		}
	}

	_, err = os.Stat(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = createConfigFile(configFilePath)
		} else {
			return "", errors.WithStack(err)
		}
	}

	return configFilePath, nil
}

func createConfigFile(configFilePath string) error {
	// load template
	statikFS, err := fs.New()
	if err != nil {
		return errors.WithStack(err)
	}
	cfgTpl, err := statikFS.Open(configTemplatePath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer cfgTpl.Close()
	b, err := ioutil.ReadAll(cfgTpl)
	if err != nil {
		return errors.WithStack(err)
	}

	// create config file from template
	err = ioutil.WriteFile(configFilePath, b, 0755)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("create %s\n", configFilePath)

	return nil
}

func getConfigPath() (string, error) {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if len(xdgConfigHome) == 0 {
		xdgConfigHome = filepath.Join(os.Getenv("HOME"), ".config")
	}
	appConfigHome := filepath.Join(xdgConfigHome, appName)
	appConfigFilePath := filepath.Join(appConfigHome, appConfigName)
	return appConfigFilePath, nil
}

func existsConfig(cfgFilePath string) error {
	_, err := os.Stat(cfgFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.WithStack(ErrAppConfigNotExists)
		} else {
			return errors.WithStack(err)
		}
	}

	return nil
}

// TODO: fix for handling
func validateConfig(cfg *Config) error {
	if len(cfg.Username) == 0 || len(cfg.GitHubAccessToken) == 0 {
		return ErrAppConfigValidation
	}
	return nil
}
