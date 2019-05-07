package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/grpclog"
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

	// for debug
	GithubAccessToken string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	baseDir := filepath.Clean(os.Getenv("BASE_DIR"))
	if _, err := os.Stat(baseDir); err != nil {
		grpclog.Infof("make %s\n", baseDir)
		fmt.Printf("make %s\n", baseDir)
		if err := os.Mkdir(baseDir, 0777); err != nil {
			grpclog.Errorln(err)
			fmt.Println(err)
		}
	}
	cacheDir := filepath.Join(baseDir, ".cache/")
	if _, err := os.Stat(cacheDir); err != nil {
		grpclog.Infof("make cache directory: %s\n", cacheDir)
		fmt.Printf("make cache directory: %s\n", cacheDir)
		if err := os.Mkdir(cacheDir, 0744); err != nil {
			grpclog.Errorln(err)
			fmt.Println(err)
		}
	}

	config := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		BaseDir:     baseDir,
		CacheDir:    cacheDir,

		SlackWebHookURL:    os.Getenv("SLACK_WEBHOOK_URL"),
		SlackReportChannel: "#" + os.Getenv("SLACK_REPORT_CHANNEL"),
		SlackBotName:       os.Getenv("SLACK_BOT_NAME"),
		SlackBotIcon:       ":" + os.Getenv("SLACK_BOT_ICON") + ":",

		GithubAccessToken: os.Getenv("GITHUB_ACCESS_TOKEN"),
	}

	return config, nil
}
