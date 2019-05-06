package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL     string
	SlackWebHookURL string
	CacheDir        string

	// for debug
	GithubAccessToken string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		SlackWebHookURL:   os.Getenv("SLACK_WEBHOOK_URL"),
		CacheDir:          os.Getenv("CACHE_DIR"),
		GithubAccessToken: os.Getenv("GITHUB_ACCESS_TOKEN"),
	}

	return config, nil
}
