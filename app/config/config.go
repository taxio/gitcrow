package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL string
	CacheDir    string

	// slack
	SlackWebHookURL    string
	SlackReportChannel string
	SlackBotName       string
	SlackBotIcon       string

	// for debug
	GithubAccessToken string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		CacheDir:    os.Getenv("CACHE_DIR"),

		SlackWebHookURL:    os.Getenv("SLACK_WEBHOOK_URL"),
		SlackReportChannel: "#" + os.Getenv("SLACK_REPORT_CHANNEL"),
		SlackBotName:       os.Getenv("SLACK_BOT_NAME"),
		SlackBotIcon:       ":" + os.Getenv("SLACK_BOT_ICON") + ":",

		GithubAccessToken: os.Getenv("GITHUB_ACCESS_TOKEN"),
	}

	return config, nil
}
