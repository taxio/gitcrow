package cmd

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/afero"
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

func Test_downloadManagerImpl_readCsv(t *testing.T) {
	dm := downloadManagerImpl{fs: afero.NewMemMapFs()}

	// no csv case
	_, err := dm.readCsv("download.csv")
	if err == nil {
		t.Fatalf("readCsv must return error when there is no csv")
	}

	// generate csv
	err = dm.GenerateCsv()
	if err != nil {
		t.Fatalf("dm.GenerateCsv() return error: %+v\n", err)
	}

	// read from relative path
	_, err = dm.readCsv("download.csv")
	if err != nil {
		t.Fatalf("dm.readCsv() from relative path return error: %+v\n", err)
	}

	// read from absolute path
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() return error: %+v\n", err)
	}
	p := filepath.Join(wd, "download.csv")
	_, err = dm.readCsv(p)
	if err != nil {
		t.Fatalf("dm.readCsv() from absolute path return error: %+v\n", err)
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
			name: "validation error: no data",
			fields: fields{
				fs: afero.NewMemMapFs(),
			},
			args: args{
				csvData: [][]string{
					{"owner", "repo", "tag"},
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
					{"owner", "repo", "tag"},
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

//func Test_downloadManagerImpl_send(t *testing.T) {
//	// for mock http request(RoundTripper)
//	httpmock.Activate()
//	defer httpmock.DeactivateAndReset()
//
//	dm := downloadManagerImpl{fs: afero.NewMemMapFs()}
//
//	type args struct {
//		data DownloadRequest
//		host string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *http.Response
//		wantErr bool
//	}{
//		{
//			name: "",
//			args: args{
//				data: DownloadRequest{
//					Username:    "",
//					AccessToken: "",
//					ProjectName: "",
//					Repos:       nil,
//				},
//				host: "",
//			},
//			want:    nil,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := dm.send(tt.args.data, tt.args.host)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("downloadManagerImpl.send() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("downloadManagerImpl.send() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
