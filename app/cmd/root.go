package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/log"
	"github.com/taxio/gitcrow/pkg"
)

func NewRootCmd(name, version string, subCmds []*cobra.Command) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   name,
		Short: "a tool for cloning git repositories",
		Long:  "a tool for cloning git repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			v, err := cmd.Flags().GetBool("version")
			if err != nil {
				return err
			}
			if v {
				pkg.PrintVersion(os.Stdout, version)
				return nil
			}
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			verbose, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				return err
			}
			if verbose {
				log.L().SetVerbose(true)
				log.L().Println("log configured")
			}
			return nil
		},
	}

	rootCmd.Flags().BoolP("version", "v", false, "print version")
	rootCmd.PersistentFlags().Bool("verbose", false, "print log for developer")

	// apply sub commands
	rootCmd.AddCommand(subCmds...)

	return rootCmd
}
