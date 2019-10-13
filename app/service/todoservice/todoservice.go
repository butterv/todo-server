package todoservice

import (
	"context"

	pb "github.com/istsh/todo-server/app/proto/v1/todo"
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

func (*TodoService) CreateTodo(context.Context, *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	return nil, nil
}

func (*TodoService) UpdateTodo(context.Context, *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	return nil, nil
}

func (*TodoService) DeleteTodo(context.Context, *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	return nil, nil
}

func (*TodoService) GetTodos(context.Context, *pb.GetTodosRequest) (*pb.GetTodosResponse, error) {
	return nil, nil
}
