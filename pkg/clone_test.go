package pkg

import (
	"testing"

	"github.com/spf13/afero"
)

func TestValidateRepoPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "none",
			args: args{
				path: "",
			},
			wantErr: true,
		},
		{
			name: "github",
			args: args{
				path: "github/taxio/gitcrow",
			},
			wantErr: false,
		},
		{
			name: "github with link",
			args: args{
				path: "https://github.com/taxio/gitcrow",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateRepoPath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("validateRepoPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GetLink(t *testing.T) {
	type fields struct {
		Host  string
		Owner string
		Name  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "github",
			fields: fields{
				Host:  "github.com",
				Owner: "taxio",
				Name:  "gitcrow",
			},
			want:    "https://github.com/taxio/gitcrow",
			wantErr: false,
		},
		{
			name: "missing",
			fields: fields{
				Host:  "",
				Owner: "taxio",
				Name:  "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				Host:  tt.fields.Host,
				Owner: tt.fields.Owner,
				Name:  tt.fields.Name,
			}
			got, err := r.GetLink()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneRepo(t *testing.T) {
	type args struct {
		fs           afero.Fs
		repoPath     string
		projBasePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "regular",
			args: args{
				fs:           afero.NewMemMapFs(),
				repoPath:     "github.com/taxio/gitcrow",
				projBasePath: "~/.ghq/github.com/taxio/gitcrow/.gitcrow",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CloneRepo(tt.args.fs, tt.args.repoPath, tt.args.projBasePath); (err != nil) != tt.wantErr {
				t.Errorf("CloneRepo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cloneRepo(t *testing.T) {
	type args struct {
		fs       afero.Fs
		repoPath string
		projPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				fs:       nil,
				repoPath: "",
				projPath: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cloneRepo(tt.args.fs, tt.args.repoPath, tt.args.projPath); (err != nil) != tt.wantErr {
				t.Errorf("cloneRepo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
