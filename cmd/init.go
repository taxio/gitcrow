package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/pkg"
)

func NewInitCmd(ctx *pkg.AppContext) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init gitcrow project directory",
		Long:  "init gitcrow project directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, _ = fmt.Fprintln(ctx.Out, "Not implemented yet.")
			return nil
		},
	}
	return initCmd
}
