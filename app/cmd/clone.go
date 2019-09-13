package cmd

import (
	"errors"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/pkg"
	"github.com/taxio/gitcrow/pkg/record"
)

func NewCloneCmd(r *record.RecordStore) *cobra.Command {
	cloneCmd := &cobra.Command{
		Use:   "clone",
		Short: "clone repositories",
		Long:  "clone repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("argument incorrect")
			}
			repoPath := args[0]
			projPath, err := os.Getwd()
			if err != nil {
				return err
			}

			fs := afero.NewOsFs()
			err = pkg.CloneRepo(fs, repoPath, projPath)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cloneCmd
}
