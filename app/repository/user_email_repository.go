package repository

import "github.com/istsh/todo-sample/app/domain/model"

// UserEmailRepositoryRefer is a readonly repository for user emails.
type UserEmailRepositoryRefer interface {
	FindByUserID(userID model.UserID) (*model.UserEmail, error)
	FindByEmail(email string) (*model.UserEmail, error)
}

// UserEmailRepositoryModify is a read/write repository for user emails.
type UserEmailRepositoryModify interface {
	UserEmailRepositoryRefer

	Create(userID model.UserID, email string) (*model.UserEmail, error)
	UpdateEmailByUserID(userID model.UserID, email string) error
	DeleteByUserID(userID model.UserID) error
}
