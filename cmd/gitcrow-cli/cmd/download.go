package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Request repositories download",
	Long:  `Request repositories download`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringP("generate-template", "g", ".", "generate a repository list file template")
}
