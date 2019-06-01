package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var verbose = false

var rootCmd = &cobra.Command{
	Use:   "gitcrow",
	Short: "Repository downloader",
	Long:  "Repository downloader",
	Run: func(cmd *cobra.Command, args []string) {
		// verbose
		vb, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			log.Fatal(err)
		}
		verbose = vb
		if verbose {
			log.Println("--- Run as Debug Mode ---")
		}

		// version
		v, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatal(err)
		}
		if v {
			fmt.Println("gitcrow-cli v0.0.1")
			return
		}

		// ping
		p, err := cmd.Flags().GetBool("ping")
		if err != nil {
			log.Fatal(err)
		}
		if p {
			log.Fatal("not implemented yet")
		}
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.Flags().Bool("verbose", false, "show debug log")
	rootCmd.Flags().BoolP("ping", "p", false, "send a ping to the server")
}

func Execute() error {
	return rootCmd.Execute()
}
