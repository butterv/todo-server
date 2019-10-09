package model

import "time"

type UserID string

type User struct {
	id        UserID
	createdAt time.Time
}

func (u *User) GetID() UserID {
	return u.id
}
