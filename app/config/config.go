package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type Config struct {
	DatabaseURL string
	BaseDir     string
	CacheDir    string

	// slack
	SlackWebHookURL    string
	SlackReportChannel string
	SlackBotName       string
	SlackBotIcon       string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.WithStack(err)
	}

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

	// slack configuration
	slackWebHookURL, err := getRequiredEnv("SLACK_WEBHOOK_URL")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	slackReportChannel, err := getRequiredEnv("SLACK_REPORT_CHANNEL")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	slackBotName, err := getRequiredEnv("SLACK_BOT_NAME")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	slackBotIcon, err := getRequiredEnv("SLACK_BOT_ICON")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	config := &Config{
		DatabaseURL: databaseURL,
		BaseDir:     baseDir,
		CacheDir:    cacheDir,

		SlackWebHookURL:    slackWebHookURL,
		SlackReportChannel: "#" + slackReportChannel,
		SlackBotName:       slackBotName,
		SlackBotIcon:       ":" + slackBotIcon + ":",
	}

	return config, nil
}

func getRequiredEnv(envName string) (string, error) {
	env := os.Getenv(envName)
	if len(env) == 0 {
		return "", errors.New(fmt.Sprintf("[.env] %s is required", envName))
	}

	return env, nil
}
