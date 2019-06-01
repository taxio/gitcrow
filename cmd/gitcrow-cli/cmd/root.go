package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitcrow",
	Short: "Repository downloader",
	Long:  "Repository downloader",
	Run: func(cmd *cobra.Command, args []string) {
		// version
		v, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatal(err)
		}
		if v {
			fmt.Println("gitcrow-cli v0.0.1")
			return
		}
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.Flags().Bool("verbose", false, "show debug log")
}

func Execute() error {
	return rootCmd.Execute()
}
