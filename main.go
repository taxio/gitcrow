package main

import (
	"log"
	"os"

	"github.com/spf13/afero"
	"github.com/taxio/gitcrow/cmd"
	"github.com/taxio/gitcrow/pkg"
)

const (
	name    = "gitcrow"
	version = "v0.0.1"
)

func main() {
	appCtx := pkg.AppContext{
		Name:    name,
		Version: version,
		Fs:      afero.NewOsFs(),
		Out:     os.Stdout,
	}

	cli := cmd.NewRootCmd(&appCtx)
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
