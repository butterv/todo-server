package repository

import "github.com/istsh/todo-server/app/domain/model"

// UserPasswordRepositoryRefer is a readonly repository for user passwords.
type UserPasswordRepositoryRefer interface {
	FindByUserID(userID model.UserID) (*model.UserPassword, error)
}

// UserPasswordRepositoryModify is a read/write repository for user passwords.
type UserPasswordRepositoryModify interface {
	UserPasswordRepositoryRefer

	Create(userID model.UserID, passwordHash string) (*model.UserPassword, error)
	UpdatePasswordHashByUserID(userID model.UserID, passwordHash string) error
	DeleteByUserID(userID model.UserID) error
}
