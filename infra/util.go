package infra

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/grpclog"
)

var ErrInvalidFilePath = errors.New("invalid file path")

func ValidateUserFilePath(ctx context.Context, username, projectName, filename string) error {
	if strings.Index(username, "..") != -1 {
		return errors.WithStack(ErrInvalidFilePath)
	}
	if strings.Index(projectName, "..") != -1 {
		return errors.WithStack(ErrInvalidFilePath)
	}
	if strings.Index(filename, "..") != -1 {
		return errors.WithStack(ErrInvalidFilePath)
	}
	p := filepath.Join(username, projectName, filename)
	p = filepath.Clean(p)
	ps := strings.Split(p, string(os.PathSeparator))
	if len(ps) != 3 {
		return errors.WithStack(ErrInvalidFilePath)
	}
	return nil
}

func MkdirRecurrently(ctx context.Context, baseDir, username, projectName string) error {
	// make user dir if not exist
	p := filepath.Join(baseDir, username)
	if _, err := os.Stat(p); err != nil {
		grpclog.Infof("make dir: %s\n", p)
		if err := os.Mkdir(p, 0777); err != nil {
			return errors.WithStack(err)
		}
	}

	// make project dir if not exist
	p = filepath.Join(p, projectName)
	if _, err := os.Stat(p); err != nil {
		grpclog.Infof("make dir: %s\n", p)
		if err := os.Mkdir(p, 0777); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
