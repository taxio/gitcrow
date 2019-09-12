package main

import (
	"os"

	"github.com/taxio/gitcrow/app"

	_ "github.com/mattn/go-sqlite3"
	"github.com/taxio/gitcrow/log"
)

const (
	name    = "gitcrow"
	version = "v0.0.1"
)

func main() {
	//appCtx := &pkg.AppContext{
	//	Name:    name,
	//	Version: version,
	//	Fs:      afero.NewOsFs(),
	//	Out:     os.Stdout,
	//}
	//
	//cli := cmd.NewRootCmd(appCtx)
	//if err := cli.Execute(); err != nil {
	//	log.L().Error(err)
	//	os.Exit(1)
	//}
	//os.Exit(0)
	err := app.Run()
	if err != nil {
		log.L().Error(err)
		os.Exit(1)
	}
	os.Exit(0)
}
