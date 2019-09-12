package pkg

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/taxio/gitcrow/db"
)

func InitProject(fs afero.Fs, projPath string) error {
	// create gitcrow dir to the project dir
	cfgPath, err := createConfigDir(fs, projPath)

	// create sqlite3 db
	dbPath := filepath.Join(cfgPath, "db.sqlite3")
	_, err = db.CreateDatabase(dbPath)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	return nil
}

func createConfigDir(fs afero.Fs, projPath string) (string, error) {
	af := afero.Afero{Fs: fs}
	cfgPath := filepath.Join(projPath, ".gitcrow")
	ext, err := af.DirExists(cfgPath)
	if err != nil {
		return "", err
	}
	if ext {
		return "", errors.New(fmt.Sprintf("%s already exists", cfgPath))
	}
	err = af.Mkdir(cfgPath, 0766)
	if err != nil {
		return "", err
	}
	return cfgPath, nil
}
