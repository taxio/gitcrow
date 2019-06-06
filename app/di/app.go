package di

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/taxio/gitcrow/app/config"
	"github.com/taxio/gitcrow/domain/repository"
	"github.com/taxio/gitcrow/infra"
)

type AppComponent interface {
	CacheStore() repository.CacheStore
	RecordStore() repository.RecordStore
	ReportStore() repository.ReportStore
	UserStore() repository.UserStore
}

func CreateAppComponent(cfg *config.Config) (AppComponent, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &appComponentImpl{
		config: cfg,
		db:     db,
	}, nil
}

type appComponentImpl struct {
	config *config.Config
	db     *sql.DB
}

func (c *appComponentImpl) CacheStore() repository.CacheStore {
	return infra.NewCacheStore(c.config.CacheDir)
}

func (c *appComponentImpl) RecordStore() repository.RecordStore {
	return infra.NewRecordStore(c.db)
}

func (c *appComponentImpl) ReportStore() repository.ReportStore {
	return infra.NewReportStore(
		c.config.Xhook.Url,
		c.config.Xhook.Channel,
		c.config.Xhook.BotName,
		c.config.Xhook.BotIcon,
		c.config.BaseDir,
	)
}

func (c *appComponentImpl) UserStore() repository.UserStore {
	return infra.NewUserStore(c.config.BaseDir)
}
