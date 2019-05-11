package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"

	"github.com/pkg/errors"
)

const version = "v0.0.1a1"

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Printf("GitCrow CLI %s\n", version)

	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	pp.Println(cfg)

	return nil
}

func getConfigDir() (string, error) {
	// Load env
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if len(configDir) == 0 {
		configDir = filepath.Join(os.Getenv("HOME"), ".config")
	}

	_, err := os.Stat(configDir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.Errorf("please config base directory: %s", configDir)
		} else {
			return "", errors.WithStack(err)
		}
	}

	configDir = filepath.Join(configDir, "gitcrow")

	return configDir, nil
}

type Config struct {
	Username          string
	GitHubAccessToken string
}

func loadConfig() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	// config dir存在確認 & 無いなら作る
	_, err = os.Stat(configDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(configDir, 0777)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		} else {
			return nil, errors.WithStack(err)
		}
	}

	configFile := filepath.Join(configDir, "gitcrow.toml")
	_, err = os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			err = createConfigFile(configFile)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		} else {
			return nil, errors.WithStack(err)
		}
	}

	return nil, nil
}

func createConfigFile(filename string) error {
	// TODO: load template
	// TODO: create config file from template
	// TODO: print message
	return nil
}
