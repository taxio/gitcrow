package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL     string
	SlackWebHookURL string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		SlackWebHookURL: os.Getenv("SLACK_WEBHOOK_URL"),
	}

	return config, nil
}
