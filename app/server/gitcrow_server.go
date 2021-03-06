package server

import (
	"context"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"
	api_pb "github.com/taxio/gitcrow/api"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type GitcrowServiceServer interface {
	api_pb.GitcrowServiceServer
	grapiserver.Server
}

func NewGitcrowServiceServer(downloadService service.DownloadService, cloneService service.CloneService) GitcrowServiceServer {
	return &gitcrowServiceServerImpl{
		downloadSvc: downloadService,
		cloneSvc:    cloneService,
	}
}

type gitcrowServiceServerImpl struct {
	downloadSvc service.DownloadService
	cloneSvc    service.CloneService
}

func (s *gitcrowServiceServerImpl) CloneRepositories(ctx context.Context, req *api_pb.CloneRepositoriesRequest) (*api_pb.CloneRepositoriesResponse, error) {
	// TODO: Not yet implemented
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (s *gitcrowServiceServerImpl) DownloadRepositories(ctx context.Context, req *api_pb.DownloadRepositoriesRequest) (*api_pb.DownloadRepositoriesResponse, error) {
	grpclog.Infof("request by %s\n", req.Username)

	var repos []*model.GitRepo
	for _, repo := range req.Repos {
		repos = append(repos, &model.GitRepo{
			Owner:   repo.Owner,
			Repo:    repo.Repo,
			Tag:     repo.GetTag().Value,
			IsClone: false,
		})
	}

	err := s.downloadSvc.DelegateToWorker(ctx, req.Username, req.ProjectName, req.AccessToken, repos)
	if err != nil {
		switch errors.Cause(err) {
		case service.ErrGitHubAuth:
			return nil, status.Error(codes.Unauthenticated, "invalid github token")
		case service.ErrPathValidation:
			return nil, status.Error(codes.InvalidArgument, "invalid path")
		case service.ErrAlreadyAcceptedDownloadRequest:
			return nil, status.Error(codes.AlreadyExists, "you already requested")
		}
		return nil, status.Error(codes.Internal, "internal error, sorry")
	}

	return &api_pb.DownloadRepositoriesResponse{Message: "request accepted"}, nil
}
