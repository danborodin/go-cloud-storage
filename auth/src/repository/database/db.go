package database

import "auth/src/types"

type DB interface {
	AddUser(user *types.User) error
}
