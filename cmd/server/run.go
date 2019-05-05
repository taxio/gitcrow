package main

import (
	"context"
	"log"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/k0kubun/pp"
	"github.com/taxio/gitcrow/app/config"
	"github.com/taxio/gitcrow/app/di"
	"github.com/taxio/gitcrow/app/server"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/service"
)

func run() error {
	// Application context
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	pp.Println(cfg)

	_, err = di.CreateAppComponent(cfg)
	if err != nil {
		return err
	}

	return nil

	downloadSvc := service.NewDownloadService()
	repos := []*model.GitRepo{
		&model.GitRepo{
			Owner:   "taxio",
			Repo:    "gitcrow1",
			Tag:     "v0.0.1",
			IsClone: false,
		},
		&model.GitRepo{
			Owner:   "taxio",
			Repo:    "gitcrow2",
			Tag:     "v0.0.2",
			IsClone: false,
		},
		&model.GitRepo{
			Owner:   "taxio",
			Repo:    "gitcrow3",
			Tag:     "v0.0.3",
			IsClone: false,
		},
	}
	err = downloadSvc.DelegateToWorker(ctx, "taxio", "/tmp", cfg.GithubAccessToken, repos)
	if err != nil {
		log.Fatal(err)
	}

	return nil
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewGitcrowServiceServer(),
		),
	)
	return s.ServeContext(ctx)
}
