package model

// UserEmail is a structure with user email information.
type UserEmail struct {
	userID UserID
	email  string
}

// NewUserEmail initializes UserEmail.
func NewUserEmail(userID UserID, email string) *UserEmail {
	return &UserEmail{
		userID: userID,
		email:  email,
	}
}

// GetUserID gets a user id in UserEmail.
func (u *UserEmail) GetUserID() UserID {
	return u.userID
}

// GetEmail gets a email in UserEmail.
func (u *UserEmail) GetEmail() string {
	return u.email
}
