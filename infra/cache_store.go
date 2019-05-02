package infra

import (
	"context"
	"github.com/taxio/gitcrow/domain/model"
	"io"
)

type cacheStoreImpl struct {
	SaveDir string
}

func NewCacheStore(saveDir string) *cacheStoreImpl {
	return &cacheStoreImpl{
		SaveDir: saveDir,
	}
}

func (s *cacheStoreImpl) Exists(ctx context.Context, repo model.GitRepo) (bool, error) {
	return false, nil
}

func (s *cacheStoreImpl) Save(ctx context.Context, filename string, data io.ReadCloser) error {
	return nil
}
