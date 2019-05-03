package repository

import (
	"context"
	"github.com/taxio/gitcrow/domain/model"
	"io"
)

type CacheStore interface {
	Exists(ctx context.Context, repo model.GitRepo) (bool, error)
	Save(ctx context.Context, filename string, data io.ReadCloser) error
}

type RecordStore interface {
	Exists(ctx context.Context, repo model.GitRepo) (bool, error)
	Insert(ctx context.Context, repo model.GitRepo) error
	Sync(ctx context.Context, repos *[]model.GitRepo) error
}

type UserStore interface {
	Save(ctx context.Context, filename string, data io.ReadCloser) error
	Clone(ctx context.Context, repo model.GitRepo) error
}

type ReportStore interface {
	Notify(ctx context.Context, username, message string) error
	Save(ctx context.Context) error
}