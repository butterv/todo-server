package model

type UserPassword struct {
	userID       UserID
	passwordHash string
}

func (u *UserPassword) GetUserID() UserID {
	return u.userID
}

func (u *UserPassword) GetPasswordHash() string {
	return u.passwordHash
}
