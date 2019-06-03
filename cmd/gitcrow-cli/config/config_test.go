package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/afero"
	_ "github.com/taxio/gitcrow/cmd/gitcrow-cli/statik"
)

func TestNewManager(t *testing.T) {
	type args struct {
		fs afero.Fs
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "MemMapFs",
			args: args{
				afero.NewMemMapFs(),
			},
		},
		{
			name: "OsFs",
			args: args{
				afero.NewOsFs(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewManager(tt.args.fs); reflect.ValueOf(got).IsNil() {
				t.Error("NewManager() returns nil")
			}
		})
	}
}

func TestManagerImpl_Exists(t *testing.T) {
	fs := afero.NewMemMapFs()
	cm := NewManager(fs)

	got, err := cm.Exists()
	if err != nil {
		t.Fatalf("Exists() returns err: %+v", err)
	}
	if got {
		t.Fatal("first Exists() must return false, but got true")
	}

	// create config dir and file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	appConfigDir := filepath.Join(homeDir, ".config", "gitcrow")
	err = fs.MkdirAll(appConfigDir, 0744)
	if err != nil {
		t.Fatal(err)
	}
	appConfigFilePath := filepath.Join(appConfigDir, "config.toml")
	_, err = fs.Create(appConfigFilePath)
	if err != nil {
		t.Fatal(err)
	}

	// check config exists
	got, err = cm.Exists()
	if err != nil {
		t.Fatalf("Exists() returns err: %+v", err)
	}
	if !got {
		t.Fatal("Exists() must return true, but got false")
	}
}

func TestManagerImpl_GenerateFromTemplate(t *testing.T) {
	fs := afero.NewMemMapFs()
	cm := NewManager(fs)

	err := cm.GenerateFromTemplate("", "", "")
	if err != nil {
		t.Fatal(err)
	}
	ext, err := cm.Exists()
	if err != nil {
		t.Fatal(err)
	}
	if !ext {
		t.Fatal("config file has not been created")
	}
}

func TestManagerImpl_Load(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			name: "no value",
			want: &Config{
				ServerHost:        "",
				Username:          "",
				GitHubAccessToken: "",
			},
			wantErr: false,
		},
		{
			name: "normal",
			want: &Config{
				ServerHost:        "localhost",
				Username:          "foo",
				GitHubAccessToken: "hoge",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &managerImpl{
				fs: afero.NewMemMapFs(),
			}
			err := c.GenerateFromTemplate(tt.want.ServerHost, tt.want.Username, tt.want.GitHubAccessToken)
			if err != nil {
				t.Errorf("GenerateFromTemplate() error = %+v\n", err)
				return
			}
			got, err := c.Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
