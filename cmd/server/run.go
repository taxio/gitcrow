package main

import (
	"context"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/taxio/gitcrow/app/config"
	"github.com/taxio/gitcrow/app/di"
	"github.com/taxio/gitcrow/app/server"
	"github.com/taxio/gitcrow/domain/service"
	"github.com/volatiletech/sqlboiler/boil"
)

func run() error {
	// Application context
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	appComp, err := di.CreateAppComponent(cfg)
	if err != nil {
		return err
	}

	downloadSvc := service.NewDownloadService(appComp)
	var cloneSvc service.CloneService // TODO: impl

	// sqlboiler configure
	boil.DebugMode = false

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewGitcrowServiceServer(downloadSvc, cloneSvc),
		),
	)
	return s.ServeContext(ctx)
}
