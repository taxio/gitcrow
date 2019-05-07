package server

import (
	"context"
	api_pb "github.com/taxio/gitcrow/api"
	"testing"
)

func TestGitcrowServiceServerImpl_CloneRepositories(t *testing.T) {
	svr := NewGitcrowServiceServer(nil, nil)

	ctx := context.Background()
	req := &api_pb.CloneRepositoriesRequest{}

	resp, err := svr.CloneRepositories(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func TestGitcrowServiceServerImpl_DownloadRepositories(t *testing.T) {
	svr := NewGitcrowServiceServer(nil, nil)

	ctx := context.Background()
	req := &api_pb.DownloadRepositoriesRequest{}

	resp, err := svr.DownloadRepositories(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
