package services

import (
	"auth/src/repository/database"
	"auth/src/services/email"
	"auth/src/services/email/gmail"
	"auth/src/services/user"
	"auth/src/services/vault"

	"github.com/danborodin/go-logd"
)

type Services struct {
	l            *logd.Logger
	UserService  *user.Service
	VaultService *vault.Service
	EmailService email.Service
}

func NewServices(l *logd.Logger, db database.DB) *Services {
	vaultSrvc := vault.NewVaultService(l)
	emailSrvc := gmail.NewGmailService(l, db)
	userSrvc := user.NewUserService(l, db, vaultSrvc, emailSrvc)
	return &Services{
		UserService:  userSrvc,
		VaultService: vaultSrvc,
		EmailService: emailSrvc,
	}
}
