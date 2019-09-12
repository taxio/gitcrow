package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewConfigCmd() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "show config",
		Long:  "show config",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, _ = fmt.Fprintln(ctx.Out, "Not implemented yet.")
			return nil
		},
	}

	return configCmd
}
