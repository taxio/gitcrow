package infra

import (
	"context"
	"github.com/taxio/gitcrow/domain/model"
	"io"
)

type userStoreImpl struct {}

func NewUserStore() *userStoreImpl {
	return &userStoreImpl{}
}

func (s *userStoreImpl) Save(ctx context.Context, filename string, data io.ReadCloser) error {
	return nil
}

func (s *userStoreImpl) Clone(ctx context.Context, repo model.GitRepo) error {
	return nil
}