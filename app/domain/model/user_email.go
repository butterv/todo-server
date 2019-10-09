package model

type UserEmail struct {
	userID UserID
	email  string
}

func (u *UserEmail) GetUserID() UserID {
	return u.userID
}

func (u *UserEmail) GetEmail() string {
	return u.email
}
