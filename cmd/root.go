package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewRootCmd(ctx *pkg.AppContext) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gitcrow",
		Short: "a tool for cloning git repositories",
		Long:  "a tool for cloning git repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			v, err := cmd.Flags().GetBool("version")
			if err != nil {
				return err
			}
			if v {
				pkg.PrintVersion(ctx.Out, ctx.Version)
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
				err = initLogger()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	rootCmd.Flags().BoolP("version", "v", false, "print version")
	rootCmd.PersistentFlags().Bool("verbose", false, "print log for developer")

	// sub commands
	subCmds := []*cobra.Command{
		NewInitCmd(ctx),
		NewCloneCmd(ctx),
		NewDownloadCmd(ctx),
		NewConfigCmd(ctx),
	}
	rootCmd.AddCommand(subCmds...)

	return rootCmd
}

func initLogger() error {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	zapLogger, err := zapCfg.Build()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(zapLogger)
	return nil
}
