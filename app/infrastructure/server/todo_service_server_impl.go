package server

import (
	"context"

	"github.com/istsh/todo-server/app/entity/repository"
	todopb "github.com/istsh/todo-server/app/pb/v1/todo"
	"github.com/istsh/todo-server/app/usecase"
)

type todoServiceServer struct {
	r repository.Repository
	u usecase.UserUserCase
}

// NewLoginTodoServiceServer creates todo service server implementation.
func NewTodoServiceServer(r repository.Repository, u usecase.UserUserCase) todopb.TodoServiceServer {
	return &todoServiceServer{
		r: r,
		u: u,
	}
}

func (s *todoServiceServer) Authenticate(ctx context.Context, req interface{}) (context.Context, error) {
	return ctx, nil
}

func (s *todoServiceServer) CreateTodo(context.Context, *todopb.CreateTodoRequest) (*todopb.CreateTodoResponse, error) {
	return nil, nil
}

func (s *todoServiceServer) UpdateTodo(context.Context, *todopb.UpdateTodoRequest) (*todopb.UpdateTodoResponse, error) {
	return nil, nil
}

func (s *todoServiceServer) DeleteTodo(context.Context, *todopb.DeleteTodoRequest) (*todopb.DeleteTodoResponse, error) {
	return nil, nil
}

func (s *todoServiceServer) GetTodos(context.Context, *todopb.GetTodosRequest) (*todopb.GetTodosResponse, error) {
	return nil, nil
}
