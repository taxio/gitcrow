package di

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/app/cmd"
	"github.com/taxio/gitcrow/pkg/record"
)

func provideRootCmd(appCtx *AppContext, subCmds []*cobra.Command) *cobra.Command {
	return cmd.NewRootCmd(appCtx.Name, appCtx.Version, subCmds)
}

func provideSubCmds(appCtx *AppContext, r *record.RecordStore) []*cobra.Command {
	downloadCmd := cmd.NewDownloadCmd()
	initCmd := cmd.NewInitCmd()
	cloneCmd := cmd.NewCloneCmd(r)
	configCmd := cmd.NewConfigCmd()

	subCmds := []*cobra.Command{
		downloadCmd,
		initCmd,
		cloneCmd,
		configCmd,
	}
	return subCmds
}

func provideRecordStore(appCtx *AppContext) (*record.RecordStore, error) {
	return record.NewRecordStore(appCtx.Fs, appCtx.DBPath), nil
}

var ProvideSet = wire.NewSet(
	provideAppContext,
	provideRecordStore,
	provideSubCmds,
	provideRootCmd,
)
