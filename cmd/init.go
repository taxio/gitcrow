package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/pkg"
	"golang.org/x/xerrors"
)

func NewInitCmd(ctx *pkg.AppContext) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init gitcrow project directory",
		Long:  "init gitcrow project directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			err = pkg.InitProject(ctx, wd)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			return nil
		},
	}
	return initCmd
}
