package main

import (
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/afero"
	"github.com/taxio/gitcrow/cmd"
	"github.com/taxio/gitcrow/log"
	"github.com/taxio/gitcrow/pkg"
)

const (
	name    = "gitcrow"
	version = "v0.0.1"
)

func main() {
	appCtx := &pkg.AppContext{
		Name:    name,
		Version: version,
		Fs:      afero.NewOsFs(),
		Out:     os.Stdout,
	}

	cli := cmd.NewRootCmd(appCtx)
	if err := cli.Execute(); err != nil {
		log.L().Error(err)
		os.Exit(1)
	}
	os.Exit(0)
}
