package di

import (
	"io"

	"github.com/spf13/afero"
)

var (
	appName    = "gitcrow"
	appVersion = "v0.0.1"
)

type AppContext struct {
	Name    string
	Version string

	Fs  afero.Fs
	Out io.Writer
}

func provideAppContext() *AppContext {
	return &AppContext{
		Name:    appName,
		Version: appVersion,
	}
}
