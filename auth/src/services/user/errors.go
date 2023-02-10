package user

import "errors"

var (
	ErrInvalidEmailAddr          = errors.New("invalid email address")
	ErrEmailAlreadyUsed          = errors.New("email already used")
	ErrUsernameAlreadyUsed       = errors.New("username already used")
	ErrInvalidUsernameLength     = errors.New("invalid username length")
	ErrIncorrectVerificationCode = errors.New("incorrect verification code")
)
