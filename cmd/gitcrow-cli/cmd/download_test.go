package cmd

import (
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/afero"
	"github.com/taxio/gitcrow/cmd/gitcrow-cli/config"
	_ "github.com/taxio/gitcrow/cmd/gitcrow-cli/statik"
)

func TestNewDownloadManager(t *testing.T) {
	type args struct {
		fs afero.Fs
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "OsFs",
			args: args{
				fs: afero.NewOsFs(),
			},
		},
		{
			name: "MemMapFs",
			args: args{
				fs: afero.NewMemMapFs(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDownloadManager(tt.args.fs); got == nil {
				t.Errorf("NewDownloadManager() = %v, want nil", got)
			}
		})
	}
}

func Test_downloadManagerImpl_GenerateCsv(t *testing.T) {
	fs := afero.NewMemMapFs()
	dm := NewDownloadManager(fs)
	af := afero.Afero{Fs: fs}

	// generate csv
	err := dm.GenerateCsv()
	if err != nil {
		t.Fatalf("%+v\n", err)
	}

	// check generated file exists
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	filePath := filepath.Join(wd, "download.csv")
	ext, err := af.Exists(filePath)
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	if !ext {
		t.Fatal("download.csv not found")
	}

	// check file string
	data, err := af.ReadFile(filePath)
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	got := string(data)
	want := "owner,repository,tag"
	if got != want {
		t.Fatalf("generated csv is not correct. want: %s, got: %s", want, got)
	}
}

func Test_downloadManagerImpl_SendRequest(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		cfg     *config.Config
		csvPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				fs: afero.NewMemMapFs(),
			},
			args: args{
				cfg: &config.Config{
					ServerHost:        "",
					Username:          "",
					GitHubAccessToken: "",
				},
				csvPath: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &downloadManagerImpl{
				fs: tt.fields.fs,
			}
			if err := m.SendRequest(tt.args.cfg, tt.args.csvPath); (err != nil) != tt.wantErr {
				t.Errorf("downloadManagerImpl.SendRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_downloadManagerImpl_readCsv(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		csvPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &downloadManagerImpl{
				fs: tt.fields.fs,
			}
			got, err := m.readCsv(tt.args.csvPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("downloadManagerImpl.readCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downloadManagerImpl.readCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downloadManagerImpl_parseCsv(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		csvData [][]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []DownloadRequestRepo
		wantErr bool
	}{
		{
			name: "validation error: nil data",
			fields: fields{
				fs: afero.NewMemMapFs(),
			},
			args: args{
				csvData: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "validation error: column not correct",
			fields: fields{
				fs: afero.NewMemMapFs(),
			},
			args: args{
				csvData: [][]string{
					{"hoge", "foo"},
					{"bar", "piyo"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "normal",
			fields: fields{
				fs: afero.NewMemMapFs(),
			},
			args: args{
				csvData: [][]string{
					{"taxio", "gitcrow", "v0.0.1"},
					{"taxio2", "gitcrow2", "v0.0.2"},
					{"taxio3", "gitcrow3", "v0.1.1"},
				},
			},
			want: []DownloadRequestRepo{
				{"taxio", "gitcrow", "v0.0.1"},
				{"taxio2", "gitcrow2", "v0.0.2"},
				{"taxio3", "gitcrow3", "v0.1.1"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &downloadManagerImpl{
				fs: tt.fields.fs,
			}
			got, err := m.parseCsv(tt.args.csvData)
			if (err != nil) != tt.wantErr {
				t.Errorf("downloadManagerImpl.parseCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downloadManagerImpl.parseCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downloadManagerImpl_send(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		data DownloadRequest
		host string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &downloadManagerImpl{
				fs: tt.fields.fs,
			}
			got, err := m.send(tt.args.data, tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("downloadManagerImpl.send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downloadManagerImpl.send() = %v, want %v", got, tt.want)
			}
		})
	}
}
