package infra

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/domain/repository"
	"google.golang.org/grpc/grpclog"
)

type cacheStoreImpl struct {
	cacheDir string
	mu       *sync.Mutex
}

func NewCacheStore(cacheDir string) repository.CacheStore {
	return &cacheStoreImpl{cacheDir: cacheDir}
}

func (s *cacheStoreImpl) LoadZip(ctx context.Context, filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath.Join(s.cacheDir, filename))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return data, nil
}

func (s *cacheStoreImpl) Exists(ctx context.Context, filename string) (bool, error) {
	filename = filepath.Join(s.cacheDir, filename)
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, errors.WithStack(err)
		}
	}
	return true, nil
}

func (s *cacheStoreImpl) Save(ctx context.Context, filename string, data []byte) error {
	// create path
	p := filepath.Join(s.cacheDir, filename)
	grpclog.Infof("save to cache: %s\n", p)

	err := ioutil.WriteFile(p, data, 0644)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
