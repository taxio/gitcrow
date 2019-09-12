package app

import (
	"github.com/taxio/gitcrow/app/di"
)

func Run() error {
	app, err := di.NewApp()
	if err != nil {
		return err
	}
	err = app.Cmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
