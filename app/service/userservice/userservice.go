package userservice

import (
	"context"

	pb2 "github.com/istsh/todo-server/app/proto/v1/pb"

	"github.com/istsh/todo-server/app/repository"
)

// UserService provides a set of RPC methods.
type UserService struct {
	r repository.Repository
}

// New initializes a new UserService from Repository.
func New(r repository.Repository) *UserService {
	return &UserService{
		r: r,
	}
}

func (*UserService) CreateUser(context.Context, *pb2.CreateUserRequest) (*pb2.CreateUserResponse, error) {
	return nil, nil
}

func (*UserService) ChangeEmail(context.Context, *pb2.ChangeEmailRequest) (*pb2.ChangeEmailResponse, error) {
	return nil, nil
}

func (*UserService) ChangePassword(context.Context, *pb2.ChangePasswordRequest) (*pb2.ChangePasswordResponse, error) {
	return nil, nil
}
