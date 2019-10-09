package userservice

import (
	"context"

	pb "github.com/istsh/todo-sample/app/proto/v1/user"
	"github.com/istsh/todo-sample/app/repository"
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

func (*UserService) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, nil
}

func (*UserService) ChangeEmail(context.Context, *pb.ChangeEmailRequest) (*pb.ChangeEmailResponse, error) {
	return nil, nil
}

func (*UserService) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	return nil, nil
}
