package infra

import (
	"context"
	"database/sql"

	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
)

type recordStoreImpl struct {
	db *sql.DB
}

func NewRecordStore(db *sql.DB) repository.RecordStore {
	return &recordStoreImpl{db: db}
}

func (s *recordStoreImpl) Exists(ctx context.Context, repo model.GitRepo) (bool, error) {
	return false, nil
}

func (s *recordStoreImpl) Insert(ctx context.Context, repo model.GitRepo) error {
	return nil
}

func (s *recordStoreImpl) Sync(ctx context.Context, repos *[]model.GitRepo) error {
	return nil
}
