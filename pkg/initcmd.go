package pkg

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
)

func InitProject(appCtx *AppContext, path string) error {

	// create gitcrow dir to the project dir
	projPath, err := createProjectDir(appCtx, path)

	// create sqlite3 db
	err = createDB(projPath)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	return nil
}

func createProjectDir(appCtx *AppContext, path string) (string, error) {
	af := afero.Afero{Fs: appCtx.Fs}
	projPath := filepath.Join(path, fmt.Sprintf(".%s", appCtx.Name))
	ext, err := af.DirExists(projPath)
	if err != nil {
		return "", fmt.Errorf(": %w", err)
	}
	if ext {
		return "", errors.New(fmt.Sprintf("%s already exists", projPath))
	}
	err = af.Mkdir(projPath, 0766)
	if err != nil {
		return "", fmt.Errorf(": %w", err)
	}
	return projPath, nil
}

func createDB(gitcrowPath string) error {
	dbPath := filepath.Join(gitcrowPath, "db.sqlite3")
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	// language=sql
	q := `create table if not exists repo(id int, owner text, name text)`

	_, err = conn.Exec(q)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	return nil
}
