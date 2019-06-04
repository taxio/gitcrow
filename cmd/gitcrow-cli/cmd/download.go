package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rakyll/statik/fs"
	"github.com/shurcooL/go/ioutil"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/taxio/gitcrow/cmd/gitcrow-cli/config"
	_ "github.com/taxio/gitcrow/cmd/gitcrow-cli/statik"
	"golang.org/x/xerrors"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download [csv_path]",
	Short: "Request repositories download",
	Long:  `Request repositories download`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dm := NewDownloadManager(afero.NewOsFs())
		gen, err := cmd.Flags().GetBool("generate-template")
		if err != nil {
			return xerrors.Errorf("cmd.Flags().GetString(\"g\"): %w", err)
		}
		if gen {
			err = dm.GenerateCsv()
			if err != nil {
				return xerrors.Errorf("generateCsv(): %w", err)
			}
			return nil
		}

		// Send download request to gitcrow server
		cm := config.NewManager(afero.NewOsFs())
		cfg, err := cm.Load()
		if err != nil {
			return xerrors.Errorf("cm.Load(): %w", err)
		}

		if len(args) != 1 {
			return xerrors.New("Argument not correct")
		}
		csvPath := args[0]
		err = dm.SendRequest(cfg, csvPath)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().BoolP("generate-template", "g", false, "generate a repository list file template")
}

type DownloadManager interface {
	GenerateCsv() error
	SendRequest(cfg *config.Config, csvPath string) error
}

func NewDownloadManager(fs afero.Fs) DownloadManager {
	return &downloadManagerImpl{fs: fs}
}

type downloadManagerImpl struct {
	fs afero.Fs
}

func (m *downloadManagerImpl) GenerateCsv() error {
	fmt.Println("generate download.csv")
	statikFs, err := fs.New()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	tplFile, err := statikFs.Open("/download.csv.tmpl")
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	wd, err := os.Getwd()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	filePath := filepath.Join(wd, "download.csv")
	err = ioutil.WriteFile(filePath, tplFile)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (m *downloadManagerImpl) SendRequest(cfg *config.Config, csvPath string) error {
	fmt.Println("send request")
	fmt.Printf("csv file path: %s\n", csvPath)
	if filepath.IsAbs(csvPath) {
		wd, err := os.Getwd()
		if err != nil {
			return xerrors.Errorf("os.Getwd(): %v", err)
		}
		csvPath = filepath.Join(wd, csvPath)
		csvPath = filepath.Clean(csvPath)
	}
	return nil
}
