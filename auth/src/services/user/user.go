package user

import (
	"auth/src/configs"
	"auth/src/repository/database"
	"auth/src/services/email"
	"auth/src/services/vault"
	"auth/src/types"
	"net/mail"
	"time"

	"github.com/danborodin/go-logd"
)

type Service struct {
	l            *logd.Logger
	db           database.DB
	vaultService *vault.Service
	//emailService *email.Service
}

func NewUserService(l *logd.Logger, db database.DB, vaultSrvc *vault.Service, emailSrvc email.Service) *Service {
	return &Service{
		l:            l,
		db:           db,
		vaultService: vaultSrvc,
		//emailService: emailSrvc,
	}
}

func (s Service) RegisterUser(user *types.User) error {

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ErrInvalidEmailAddr
	}

	exist, err := s.db.EmailExist(user.Email)
	if err != nil {
		return err
	}
	if exist {
		return ErrEmailAlreadyUsed
	}

	exist, err = s.db.UsernameExist(user.Username)
	if err != nil {
		return err
	}
	if exist {
		return ErrUsernameAlreadyUsed
	}
	if len(user.Username) < 3 || len(user.Username) > 30 {
		return ErrInvalidUsernameLength
	}

	hashedPass, salt, err := s.vaultService.EncryptPwd(user.Password)
	if err != nil {
		return err
	}
	user.HashedPassword = string(hashedPass)
	user.Salt = string(salt)
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = user.CreatedAt
	user.Verified = false

	//generate verification code
	code, err := s.vaultService.GenCode()
	if err != nil {
		return err
	}
	user.Verification.Code = code

	vTime, err := time.ParseDuration(configs.Conf.EmailVerTime)
	if err != nil {
		return err
	}
	user.Verification.ExpireAt = time.Now().Add(vTime)

	err = s.db.AddUser(*user)
	if err != nil {
		return err
	}

	return err
}

// if code is 0 that means user want another code. send another code to users email
func (s Service) VerifyUser(username string, code uint64) error {
	user, err := s.db.GetUserByUsername(username)
	if err != nil {
		return err
	}

	t := time.Now()
	if user.Verification.ExpireAt.Compare(t) == -1 {
		err = s.db.DeleteUser(user.Username)
		if err != nil {
			s.l.ErrPrintln(err)
		}
		//return a http error
		return nil
	}

	if code == 0 {
		//resend code to email
	}

	if code != user.Verification.Code {
		return ErrIncorrectVerificationCode
	}

	//update users verified field
	user.Verified = true
	err = s.db.UpdateUser(*user)
	if err != nil {
		return err
	}

	return nil
}

// this way of deleting users that do not confirm their email is not reliable, use a message queue in the future
func (s Service) ClearUnverifiedUsers() {
	s.l.InfoPrintln("starting clearing routine")

	clear := func() {
		s.l.InfoPrintln("clearing unverified users")
		users, err := s.db.GetUnverifiedUsers()
		if err != nil {
			s.l.ErrPrintln(err)
		}

		t := time.Now()
		for _, v := range users {
			if v.Verification.ExpireAt.Compare(t) == -1 {
				err = s.db.DeleteUser(v.Username)
				if err != nil {
					s.l.ErrPrintln(err)
				}
			}
		}
	}
	clear()

	ticker := time.NewTicker(time.Hour * 24)
	for {
		<-ticker.C
		clear()
	}
}
