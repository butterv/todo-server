package todoservice

import (
	"context"

	pb2 "github.com/istsh/todo-server/app/proto/v1/pb"

	"github.com/istsh/todo-server/app/repository"
)

// TodoService provides a set of RPC methods.
type TodoService struct {
	r repository.Repository
}

// New initializes a new TodoService from Repository.
func New(r repository.Repository) *TodoService {
	return &TodoService{
		r: r,
	}
}

func (*TodoService) CreateTodo(context.Context, *pb2.CreateTodoRequest) (*pb2.CreateTodoResponse, error) {
	return nil, nil
}

func (*TodoService) UpdateTodo(context.Context, *pb2.UpdateTodoRequest) (*pb2.UpdateTodoResponse, error) {
	return nil, nil
}

func (*TodoService) DeleteTodo(context.Context, *pb2.DeleteTodoRequest) (*pb2.DeleteTodoResponse, error) {
	return nil, nil
}

func (*TodoService) GetTodos(context.Context, *pb2.GetTodosRequest) (*pb2.GetTodosResponse, error) {
	return nil, nil
}
