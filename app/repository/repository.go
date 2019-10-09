package repository

import "context"

// Repository provides an abstract connection method.
type Repository interface {
	NewConnection(ctx context.Context) (Connection, error)
}

// Connection provides a set of repository reader methods.
type Connection interface {
	Close() error
	Transaction(f func(Transaction) error) error

	Todo() TodoRepositoryRefer
	User() UserRepositoryRefer
	UserEmail() UserEmailRepositoryRefer
	UserPassword() UserPasswordRepositoryRefer
}

// Transaction provides a set of repository reader/writer methods.
type Transaction interface {
	Todo() TodoRepositoryModify
	User() UserRepositoryModify
	UserEmail() UserEmailRepositoryModify
	UserPassword() UserPasswordRepositoryModify
}
