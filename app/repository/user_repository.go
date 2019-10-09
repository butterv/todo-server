package repository

import "github.com/istsh/todo-sample/app/domain/model"

// UserRepositoryRefer is a readonly repository for users.
type UserRepositoryRefer interface {
	FindByID(id model.UserID) (*model.User, error)
}

// UserRepositoryModify is a read/write repository for users.
type UserRepositoryModify interface {
	UserRepositoryRefer

	Create(id model.UserID) (*model.User, error)
	DeleteByID(id model.UserID) error
}
