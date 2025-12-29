package user

import "errors"

var (

	ErrInvalidEmail       = errors.New("invalid email")

	ErrInvalidPassword    = errors.New("invalid password, must be at least 6 characters")

	ErrEmailAlreadyUsed   = errors.New("email already used")

	ErrNotFound           = errors.New("user not found")

	ErrInvalidCredentials = errors.New("invalid credentials")

)
