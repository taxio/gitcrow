package cmd

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/cmd/gitcrow-cli/config"
	"golang.org/x/xerrors"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize gitcrow configuration",
	Long:  `initialize gitcrow configuration`,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostServer, err := cmd.Flags().GetString("h")
		if err != nil {
			return err
		}
		username, err := cmd.Flags().GetString("u")
		if err != nil {
			return err
		}
		githubAccessToken, err := cmd.Flags().GetString("t")
		if err != nil {
			return err
		}

		fs := afero.NewOsFs()
		cm := config.NewManager(fs)

		// check config existence
		ext, err := cm.Exists()
		if err != nil {
			return xerrors.Errorf("cm.Exists: %v", err)
		}
		if ext {
			fmt.Println("config file already exists")
			return nil
		}
		err = cm.GenerateFromTemplate(hostServer, username, githubAccessToken)
		if err != nil {
			return xerrors.Errorf("cm.GenerateFromTemplate: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("host-server", "s", "", "Host Server Address")
	initCmd.Flags().StringP("username", "u", "", "username")
	initCmd.Flags().StringP("github-access-token", "t", "", "GitHub access token")
}
