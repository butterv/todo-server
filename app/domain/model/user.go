package model

import "time"

// UserID is type of id in User.
type UserID string

// User is a structure with user information.
type User struct {
	id        UserID
	createdAt time.Time
}

// NewUser initializes User.
func NewUser(id UserID) *User {
	return &User{
		id: id,
	}
}

// GetID gets a id in User.
func (u *User) GetID() UserID {
	return u.id
}
