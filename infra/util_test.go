package infra

import (
	"context"
	"testing"
)

func TestValidateUserFilePath(t *testing.T) {
	ctx := context.TODO()
	type args struct {
		username    string
		projectName string
		filename    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				username:    "taxio",
				projectName: "sample",
				filename:    "hoge.zip",
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				username:    "taxio/../",
				projectName: "sample",
				filename:    "hoge.zip",
			},
			wantErr: true,
		},
		{
			name: "error",
			args: args{
				username:    "taxio/foo",
				projectName: "sample",
				filename:    "hoge.zip",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateUserFilePath(ctx, tt.args.username, tt.args.projectName, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUserFilePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
