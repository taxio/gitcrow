package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	api_pb "github.com/taxio/gitcrow/api"
	"google.golang.org/grpc"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *gitcrowServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterGitcrowServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *gitcrowServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterGitcrowServiceHandler(ctx, mux, conn)
}