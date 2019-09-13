package pkg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/taxio/gitcrow/db"
)

type Repo struct {
	Host  string
	Owner string
	Name  string
}

func (r *Repo) GetLink() (string, error) {
	if len(r.Host) == 0 || len(r.Owner) == 0 || len(r.Name) == 0 {
		return "", errors.New("incorrect repo info")
	}
	link := fmt.Sprintf("https://%s/%s/%s", r.Host, r.Owner, r.Name)
	return link, nil
}

func CloneRepo(fs afero.Fs, repoPath, projBasePath string) error {
	projPath, err := os.Getwd()
	if err != nil {
		return err
	}
	dbPath := filepath.Join(projPath, ".gitcrow", "db.sqlite3")
	r := db.NewRecordStore(dbPath)
	err = r.Add("github.com", "taxio", "gitcrow")
	if err != nil {
		return err
	}
	//err := validateRepoPath(repoPath)
	//if err != nil {
	//	return err
	//}
	//repo, err := convertPathToRepo(repoPath)
	//if err != nil {
	//	return err
	//}
	//link, err := repo.GetLink()
	//if err != nil {
	//	return err
	//}
	//log.L().Printfln("clone %s", link)
	return nil
}

func validateRepoPath(path string) error {
	if len(path) == 0 {
		return errors.New("incorrect path")
	}
	return nil
}

func convertPathToRepo(repoPath string) (*Repo, error) {
	splited := strings.Split(repoPath, "/")
	repo := &Repo{
		Host:  splited[0],
		Owner: splited[1],
		Name:  splited[2],
	}
	return repo, nil
}

func cloneRepo(fs afero.Fs, repoPath, projPath string) error {
	cmd := exec.Cmd{
		Args: []string{"ls"},
	}
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}