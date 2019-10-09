package model

// UserPassword is a structure with user password information.
type UserPassword struct {
	userID       UserID
	passwordHash string
}

// NewUserPassword initializes UserPassword.
func NewUserPassword(userID UserID, passwordHash string) *UserPassword {
	return &UserPassword{
		userID:       userID,
		passwordHash: passwordHash,
	}
}

// GetUserID gets a user id in UserPassword.
func (u *UserPassword) GetUserID() UserID {
	return u.userID
}

// GetPasswordHash gets a password hash in UserPassword.
func (u *UserPassword) GetPasswordHash() string {
	return u.passwordHash
}
