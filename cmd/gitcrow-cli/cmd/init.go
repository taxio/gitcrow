package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize gitcrow configuration",
	Long:  `initialize gitcrow configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("username", "u", "", "username")
	initCmd.Flags().StringP("github-access-token", "t", "", "GitHub access token")
}
