package infra

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"google.golang.org/grpc/grpclog"
	"os"
	"path/filepath"
)

type userStoreImpl struct {
	baseDir string
}

func NewUserStore(baseDir string) repository.UserStore {
	return &userStoreImpl{
		baseDir: baseDir,
	}
}

func (s *userStoreImpl) Save(ctx context.Context, username, projectName, filename string, data []byte) error {
	// create path
	// TODO: validate for traverse
	p := filepath.Join(s.baseDir, username, projectName, filename)
	grpclog.Infof("save: %s\n", p)

	// create file
	file, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create %s", p))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			grpclog.Errorln(err)
		}
	}()

	// write zip data
	_, err = file.Write(data)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to write %s", p))
	}

	return nil
}

func (s *userStoreImpl) Clone(ctx context.Context, repo *model.GitRepo) error {
	return nil
}
