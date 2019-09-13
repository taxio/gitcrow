package di

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

var (
	appName    = "gitcrow"
	appVersion = "v0.0.1"
)

type AppContext struct {
	Name    string
	Version string

	Fs          afero.Fs
	ProjectPath string
	ConfigPath  string
	DBPath      string
}

func provideAppContext() (*AppContext, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	projPath := wd
	configPath := filepath.Join(projPath, ".gitcrow")
	dbPath := filepath.Join(configPath, "db.sqlite3")
	return &AppContext{
		Name:        appName,
		Version:     appVersion,
		Fs:          afero.NewOsFs(),
		ProjectPath: projPath,
		ConfigPath:  configPath,
		DBPath:      dbPath,
	}, nil
}
