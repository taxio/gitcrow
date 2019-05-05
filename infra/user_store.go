package infra

import (
	"context"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"io"
)

type userStoreImpl struct {}

func NewUserStore() repository.UserStore {
	return &userStoreImpl{}
}

func (s *userStoreImpl) Save(ctx context.Context, filename string, data io.ReadCloser) error {
	return nil
}

func (s *userStoreImpl) Clone(ctx context.Context, repo *model.GitRepo) error {
	return nil
}