package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/taxio/gitcrow/api"
)

// TodoServiceServer is a composite interface of api_pb.TodoServiceServer and grapiserver.Server.
type TodoServiceServer interface {
	api_pb.TodoServiceServer
	grapiserver.Server
}

// NewTodoServiceServer creates a new TodoServiceServer instance.
func NewTodoServiceServer() TodoServiceServer {
	return &todoServiceServerImpl{}
}

type todoServiceServerImpl struct {
}

func (s *todoServiceServerImpl) ListTodos(ctx context.Context, req *api_pb.ListTodosRequest) (*api_pb.ListTodosResponse, error) {
	// TODO: Not yet implemented.
	//return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
	return &api_pb.ListTodosResponse{
		Todos: []*api_pb.Todo{
			&api_pb.Todo{TodoId: "1", Title: "hoge", Done: false},
			&api_pb.Todo{TodoId: "2", Title: "foo", Done: true},
			&api_pb.Todo{TodoId: "3", Title: "bar", Done: false},
		},
	}, nil
}

func (s *todoServiceServerImpl) GetTodo(ctx context.Context, req *api_pb.GetTodoRequest) (*api_pb.Todo, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *todoServiceServerImpl) CreateTodo(ctx context.Context, req *api_pb.CreateTodoRequest) (*api_pb.Todo, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *todoServiceServerImpl) UpdateTodo(ctx context.Context, req *api_pb.UpdateTodoRequest) (*api_pb.Todo, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *todoServiceServerImpl) DeleteTodo(ctx context.Context, req *api_pb.DeleteTodoRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
