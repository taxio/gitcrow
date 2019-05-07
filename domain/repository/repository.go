package repository

import (
	"context"
	"github.com/taxio/gitcrow/domain/model"
)

type CacheStore interface {
	Exists(ctx context.Context, repo *model.GitRepo) (bool, error)
	Save(ctx context.Context, filename string, data []byte) error
}

type RecordStore interface {
	Exists(ctx context.Context, repo *model.GitRepo) (bool, error)
	Insert(ctx context.Context, repo *model.GitRepo) error
	Sync(ctx context.Context, repos []*model.GitRepo) error
	GetSlackId(ctx context.Context, username string) (string, bool, error)
}

type UserStore interface {
	Save(ctx context.Context, filename string, data []byte) error
	Clone(ctx context.Context, repo *model.GitRepo) error
}

type ReportStore interface {
	Notify(ctx context.Context, username, message string) error
	Save(ctx context.Context) error
}