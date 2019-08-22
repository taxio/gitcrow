package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/pkg"
)

func NewDownloadCmd(ctx *pkg.AppContext) *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "download repositories from GitHub",
		Long:  "download repositories from GitHub",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, _ = fmt.Fprintln(ctx.Out, "Not implemented yet.")
			return nil
		},
	}
	return downloadCmd
}
