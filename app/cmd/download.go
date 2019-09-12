package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewDownloadCmd() *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "download repositories from GitHub",
		Long:  "download repositories from GitHub",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, _ = fmt.Fprintln(os.Stdout, "Not implemented yet.")
			return nil
		},
	}
	return downloadCmd
}
