package model

import "time"

type UserID string

type User struct {
	ID        UserID
	CreatedAt time.Time
}
