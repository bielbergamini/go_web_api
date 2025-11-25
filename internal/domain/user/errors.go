package user

import "errors"

var (
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrInvalidPassword = errors.New("password must be at least 6 characters")
	ErrEmailAlreadyUsed = errors.New("email is already in use")
)
