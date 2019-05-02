package infra

import (
	"context"
	"github.com/taxio/gitcrow/domain/model"
	"io"
)

type userStoreImpl struct {
	saveDir string
}

func NewUserStore(saveDir string) *userStoreImpl {
	return &userStoreImpl{saveDir: saveDir}
}

func (s *userStoreImpl) Save(ctx context.Context, filename string, data io.ReadCloser) error {
	return nil
}

func (s *userStoreImpl) Clone(ctx context.Context, repo model.GitRepo) error {
	return nil
}