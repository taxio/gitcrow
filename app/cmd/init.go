package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/pkg"
)

func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init gitcrow project directory",
		Long:  "init gitcrow project directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf(": %w", err)
			}
			fs := afero.NewOsFs()
			err = pkg.InitProject(fs, wd)
			if err != nil {
				return fmt.Errorf(": %w", err)
			}
			return nil
		},
	}
	return initCmd
}
