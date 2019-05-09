package infra

import (
	"context"
	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"io/ioutil"
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

func (s *userStoreImpl) ValidatePathname(ctx context.Context, username, projectName string) error {
	err := ValidateUserFilePath(ctx, username, projectName, "tmp.zip")
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *userStoreImpl) MakeUserProjectDir(ctx context.Context, username, projectName string) error {
	err := MkdirRecurrently(ctx, s.baseDir, username, projectName)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *userStoreImpl) Save(ctx context.Context, username, projectName, filename string, data []byte) error {
	// validate for traverse
	err := ValidateUserFilePath(ctx, username, projectName, filename)
	if err != nil {
		return errors.WithStack(err)
	}

	p := filepath.Join(s.baseDir, username, projectName, filename)

	err = ioutil.WriteFile(p, data, 0666)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *userStoreImpl) Clone(ctx context.Context, repo *model.GitRepo) error {
	return nil
}
