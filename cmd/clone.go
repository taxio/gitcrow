package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/pkg"
)

func NewCloneCmd(ctx *pkg.AppContext) *cobra.Command {
	cloneCmd := &cobra.Command{
		Use:   "clone",
		Short: "clone repositories",
		Long:  "clone repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, _ = fmt.Fprintln(ctx.Out, "Not implemented yet.")
			return nil
		},
	}
	return cloneCmd
}
