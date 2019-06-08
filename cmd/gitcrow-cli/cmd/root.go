package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

var verbose = false

var rootCmd = &cobra.Command{
	Use:   "gitcrow",
	Short: "Repository downloader",
	Long:  "Repository downloader",
	RunE: func(cmd *cobra.Command, args []string) error {
		// version
		v, err := cmd.Flags().GetBool("version")
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
		if v {
			fmt.Println("gitcrow-cli v0.0.1")
			return nil
		}

		// ping
		p, err := cmd.Flags().GetBool("ping")
		if err != nil {
			log.Fatal(err)
		}
		if p {
			fmt.Println("not implemented yet")
			return nil
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.Flags().BoolP("ping", "p", false, "send a ping to the server")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "show debug log")

	cobra.OnInitialize(func() {
		if !verbose {
			return
		}
		var zapCfg zap.Config
		zapCfg = zap.NewDevelopmentConfig()
		zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		zapLogger, err := zapCfg.Build()
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		zap.ReplaceGlobals(zapLogger)
	})
}

func Execute() error {
	return rootCmd.Execute()
}
