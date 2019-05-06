package infra

import (
	"context"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
)

type userStoreImpl struct {}

func NewUserStore() repository.UserStore {
	return &userStoreImpl{}
}

func (s *userStoreImpl) Save(ctx context.Context, filename string, data []byte) error {
	return nil
}

func (s *userStoreImpl) Clone(ctx context.Context, repo *model.GitRepo) error {
	return nil
}