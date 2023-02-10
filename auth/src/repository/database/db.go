package database

import "auth/src/types"

type DB interface {
	// user related

	AddUser(user types.User) error
	GetUserByUsername(username string) (*types.User, error)
	UpdateUser(user types.User) error
	DeleteUser(username string) error
	EmailExist(email string) (bool, error)
	UsernameExist(username string) (bool, error)
	CheckUserVerified(username string) (bool, error)
	GetUnverifiedUsers() ([]types.User, error)

	//
}
