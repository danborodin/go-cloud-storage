package user

import (
	"auth/src/repository/database"
	"auth/src/services/vault"
	"auth/src/types"
	"github.com/danborodin/go-logd"
)

type Service struct {
	l            *logd.Logger
	db           database.DB
	vaultService *vault.Service
}

func NewUserService(l *logd.Logger, db database.DB, vaultSrvc *vault.Service) *Service {
	return &Service{
		l:            l,
		db:           db,
		vaultService: vaultSrvc,
	}
}

func (us *Service) RegisterUser(user *types.User) error {
	hashedPass, salt, err := us.vaultService.Encrypt(user.Password)
	if err != nil {
		return err
	}
	user.HashedPassword = string(hashedPass)
	user.Salt = string(salt)
	err = us.db.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}
