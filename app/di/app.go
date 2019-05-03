package di

import (
	"database/sql"
	"github.com/taxio/gitcrow/app/config"
	"github.com/taxio/gitcrow/domain/repository"
)

type AppComponent interface {
	CacheStore() repository.CacheStore
	RecordStore() repository.RecordStore
	ReportStore() repository.ReportStore
	UserStore() repository.UserStore
}

func CreateAppComponent(cfg *config.Config) (*AppComponent, error) {
	_, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	// TODO: create record store

	// TODO: create cache store

	// TODO: create report store

	// TODO: create user store

	return nil, nil
}

type appComponentImpl struct {
	config *config.Config
	db     *sql.DB
}
