package infra

import (
	"context"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"sync"
)

type cacheStoreImpl struct {
	cacheDir string
	mu       *sync.Mutex
}

func NewCacheStore(cacheDir string) repository.CacheStore {
	return &cacheStoreImpl{cacheDir: cacheDir}
}

func (s *cacheStoreImpl) Exists(ctx context.Context, repo *model.GitRepo) (bool, error) {
	return false, nil
}

func (s *cacheStoreImpl) Save(ctx context.Context, filename string, data []byte) error {
	return nil
}
