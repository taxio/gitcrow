package server

import (
	"context"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/k0kubun/pp"
	api_pb "github.com/taxio/gitcrow/api"
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
	//// TODO: Not yet implemented
	//return nil, status.Error(codes.Unimplemented, "Not implemented")
	pp.Println(req)
	return &api_pb.CloneRepositoriesResponse{Message: "request accepted"}, nil
}

func (s *gitcrowServiceServerImpl) DownloadRepositories(ctx context.Context, req *api_pb.DownloadRepositoriesRequest) (*api_pb.DownloadRepositoriesResponse, error) {
	//// TODO: Not yet implemented
	//return nil, status.Error(codes.Unimplemented, "Not implemented")
	pp.Println(req)
	return &api_pb.DownloadRepositoriesResponse{Message: "request accepted"}, nil
}
