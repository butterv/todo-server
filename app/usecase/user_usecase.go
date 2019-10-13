package usecase

import "github.com/istsh/todo-server/app/domain/model"

// UserUsecase is an user usecase.
type UserUsecase interface {
	RegisterUser(email, password string) (*model.User, error)
	GetUser(userID model.UserID) (*model.User, error)
	IsCorrectPassword(email, password string) (bool, error)
	ChangeEmail(userID model.UserID, email string) error
	ChangePassword(userID model.UserID, currentPassword, newPassword string) error
}
