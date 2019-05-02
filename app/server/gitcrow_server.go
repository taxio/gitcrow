package server

import (
	"context"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	api_pb "github.com/taxio/gitcrow/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GitcrowServiceServer interface {
	api_pb.GitcrowServiceServer
	grapiserver.Server
}

func NewGitcrowServiceServer() GitcrowServiceServer {
	return &gitcrowServiceServerImpl{}
}

type gitcrowServiceServerImpl struct {
}

func (s *gitcrowServiceServerImpl) CloneRepositories(ctx context.Context, req *api_pb.CloneRepositoriesRequest) (*api_pb.CloneRepositoriesResponse, error) {
	// TODO: Not yet implemented
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (s *gitcrowServiceServerImpl) DownloadRepositories(ctx context.Context, req *api_pb.DownloadRepositoriesRequest) (*api_pb.DownloadRepositoriesResponse, error) {
	// TODO: Not yet implemented
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}