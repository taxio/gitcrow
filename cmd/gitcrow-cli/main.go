package main

import (
	"fmt"
	"log"
	"os"

	"github.com/taxio/gitcrow/cmd/gitcrow-cli/cmd"

	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/cmd/gitcrow-cli/config"
	_ "github.com/taxio/gitcrow/cmd/gitcrow-cli/statik"
)

const version = "v0.0.1a1"
const configFileName = "gitcrow.toml"

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	os.Exit(0)
}

func run() error {
	//fmt.Printf("GitCrow CLI %s\n", version)
	//
	//cfg, err := loadConfig()
	//if err != nil {
	//	return err
	//}
	//pp.Println(cfg)

	return nil
}

func loadConfig() (*config.Config, error) {
	cfg, err := config.Load()
	if err != nil {
		if errors.Cause(err) == config.ErrAppConfigNotExists {
			// create config file
			cfgPath, err := config.CreateConfigFile()
			if err != nil {
				return nil, err
			}
			fmt.Printf("please input your information to %s\n", cfgPath)
			os.Exit(0)
		} else {
			return nil, err
		}
	}

	return cfg, nil
}
