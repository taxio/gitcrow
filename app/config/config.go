package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	DatabaseURL string
	BaseDir     string
	CacheDir    string

	// xhook
	Xhook XhookConfig
}

type XhookConfig struct {
	Url     string
	Channel string
	BotName string
	BotIcon string
}

func Load() (*Config, error) {
	// ignore error for docker
	_ = godotenv.Load()

	databaseURL, err := getRequiredEnv("DATABASE_URL")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	baseDir := filepath.Clean(os.Getenv("BASE_DIR"))
	if _, err := os.Stat(baseDir); err != nil {
		fmt.Printf("INFO: make %s\n", baseDir)
		if err := os.Mkdir(baseDir, 0777); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	cacheDir := filepath.Join(baseDir, ".cache/")
	if _, err := os.Stat(cacheDir); err != nil {
		fmt.Printf("INFO: make cache directory: %s\n", cacheDir)
		if err := os.Mkdir(cacheDir, 0744); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// xhook configuration
	xhookCfg, err := loadXhookConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	config := &Config{
		DatabaseURL: databaseURL,
		BaseDir:     baseDir,
		CacheDir:    cacheDir,
		Xhook:       xhookCfg,
	}

	return config, nil
}

func loadXhookConfig() (XhookConfig, error) {
	cfg := XhookConfig{}

	xhookUrl, err := getRequiredEnv("XHOOK_URL")
	if err != nil {
		return cfg, errors.WithStack(err)
	}
	cfg.Url = xhookUrl

	xhookChan, err := getRequiredEnv("XHOOK_CHANNEL")
	if err != nil {
		return cfg, errors.WithStack(err)
	}
	cfg.Channel = xhookChan

	xhookBotName, err := getRequiredEnv("XHOOK_BOT_NAME")
	if err != nil {
		return cfg, errors.WithStack(err)
	}
	cfg.BotName = xhookBotName

	xhookBotIcon, err := getRequiredEnv("XHOOK_BOT_ICON")
	if err != nil {
		return cfg, errors.WithStack(err)
	}
	cfg.BotIcon = xhookBotIcon

	return cfg, nil
}

func getRequiredEnv(envName string) (string, error) {
	env := os.Getenv(envName)
	if len(env) == 0 {
		return "", errors.New(fmt.Sprintf("environment %s is required", envName))
	}

	return env, nil
}
