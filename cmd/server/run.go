package main

import (
	"context"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/taxio/gitcrow/app/config"
	"github.com/taxio/gitcrow/app/di"
	"github.com/taxio/gitcrow/app/server"
)

func run() error {
	// Application context
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	_, err = di.CreateAppComponent(cfg)
	if err != nil {
		return err
	}

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewGitcrowServiceServer(),
		),
	)
	return s.ServeContext(ctx)
}
