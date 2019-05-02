package server

import (
	"context"
	"testing"

	api_pb "github.com/taxio/gitcrow/api"
)

func Test_TodoServiceServer_ListTodos(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.ListTodosRequest{}

	resp, err := svr.ListTodos(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TodoServiceServer_GetTodo(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.GetTodoRequest{}

	resp, err := svr.GetTodo(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TodoServiceServer_CreateTodo(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateTodoRequest{}

	resp, err := svr.CreateTodo(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TodoServiceServer_UpdateTodo(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateTodoRequest{}

	resp, err := svr.UpdateTodo(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TodoServiceServer_DeleteTodo(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteTodoRequest{}

	resp, err := svr.DeleteTodo(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
