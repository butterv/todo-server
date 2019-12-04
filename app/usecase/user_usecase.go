package usecase

import (
	"context"

	"github.com/istsh/todo-server/app/entity/model"
	"github.com/istsh/todo-server/app/entity/repository"
)

type UserUserCase interface {
	CreateUser(ctx context.Context, tx repository.Transaction, email, password string) error
	IsCorrectUserPassword(ctx context.Context, con repository.Connection, userID model.UserID, password string) (bool, error)

	CreateUserToken(ctx context.Context, tx repository.Transaction, userID model.UserID) (string, error)
}
