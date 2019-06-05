package cmd

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rakyll/statik/fs"
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
	tplData, err := ioutil.ReadAll(tplFile)
	af := afero.Afero{Fs: m.fs}
	wd, err := os.Getwd()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	filePath := filepath.Join(wd, "download.csv")
	err = af.WriteFile(filePath, tplData, 0744)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (m *downloadManagerImpl) readCsv(csvPath string) ([][]string, error) {
	if filepath.IsAbs(csvPath) {
		wd, err := os.Getwd()
		if err != nil {
			return nil, xerrors.Errorf("os.Getwd(): %v", err)
		}
		csvPath = filepath.Join(wd, csvPath)
		csvPath = filepath.Clean(csvPath)
	}
	file, err := m.fs.Open(csvPath)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	reader := csv.NewReader(file)
	recs, err := reader.ReadAll()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return recs, nil
}

func (m *downloadManagerImpl) parseCsv(csvData [][]string) ([]DownloadRequestRepo, error) {
	if csvData == nil {
		return nil, xerrors.New("csv validation error. csv data is nil.")
	}
	repos := make([]DownloadRequestRepo, 0, len(csvData))
	for _, rec := range csvData {
		if len(rec) != 3 {
			return nil, xerrors.New("csv validation error. number of columns != 3")
		}
		repos = append(repos, DownloadRequestRepo{
			Owner: rec[0],
			Repo:  rec[1],
			Tag:   rec[2],
		})
	}

	return repos, nil
}

func (m *downloadManagerImpl) send(data DownloadRequest, host string) (*http.Response, error) {
	return nil, nil
}

func (m *downloadManagerImpl) SendRequest(cfg *config.Config, csvPath string) error {
	log.Println("send request")
	log.Printf("csv file path: %s\n", csvPath)

	csvData, err := m.readCsv(csvPath)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	repos, err := m.parseCsv(csvData)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	data := DownloadRequest{
		Username:    cfg.Username,
		AccessToken: cfg.GitHubAccessToken,
		ProjectName: "",
		Repos:       repos,
	}

	_, err = m.send(data, cfg.ServerHost)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

type DownloadRequest struct {
	Username    string                `json:"username"`
	AccessToken string                `json:"access_token"`
	ProjectName string                `json:"project_name"`
	Repos       []DownloadRequestRepo `json:"repos"`
}

type DownloadRequestRepo struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
	Tag   string `json:"tag"`
}
