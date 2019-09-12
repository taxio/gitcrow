package di

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/app/cmd"
)

func provideRootCmd(appCtx *AppContext, subCmds []*cobra.Command) *cobra.Command {
	return cmd.NewRootCmd(appCtx.Name, appCtx.Version, subCmds)
}

func provideSubCmds(appCtx *AppContext) []*cobra.Command {
	downloadCmd := cmd.NewDownloadCmd()
	initCmd := cmd.NewInitCmd()
	cloneCmd := cmd.NewCloneCmd()
	configCmd := cmd.NewConfigCmd()

	subCmds := []*cobra.Command{
		downloadCmd,
		initCmd,
		cloneCmd,
		configCmd,
	}
	return subCmds
}

var ProvideSet = wire.NewSet(
	provideAppContext,
	provideSubCmds,
	provideRootCmd,
)
