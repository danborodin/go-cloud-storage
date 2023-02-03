package services

import (
	"auth/src/repository/database"
	"auth/src/services/user"
	"auth/src/services/vault"

	"github.com/danborodin/go-logd"
)

type Services struct {
	l            *logd.Logger
	UserService  *user.Service
	VaultService *vault.Service
}

func NewServices(l *logd.Logger, db database.DB) *Services {
	vaultSrvc := vault.NewVaultService(l)
	userSrvc := user.NewUserService(l, db, vaultSrvc)
	return &Services{
		UserService:  userSrvc,
		VaultService: vaultSrvc,
	}
}
