package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/pkg"
)

func NewCloneCmd(ctx *pkg.AppContext) *cobra.Command {
	cloneCmd := &cobra.Command{
		Use:   "clone",
		Short: "clone repositories",
		Long:  "clone repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("argument incorrect")
			}
			repoPath := args[0]

			err := pkg.CloneRepo(ctx.Fs, repoPath, "")
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cloneCmd
}
