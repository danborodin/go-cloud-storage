package vault

import (
	"auth/src/configs"
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/danborodin/go-logd"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	l *logd.Logger
}

func NewVaultService(l *logd.Logger) *Service {
	return &Service{
		l: l,
	}
}

func (s Service) EncryptPwd(password string) ([]byte, []byte, error) {

	pepper := configs.Conf.Pepper
	salt, err := s.generateSalt()
	if err != nil {
		s.l.ErrPrintln(err)
		return nil, nil, err
	}

	// pepper is 30 bytes, password is min 10 max 22 bytes, salt is min 1 max 20 bytes, total is max 72 bytes
	newPass := pepper + password + string(salt)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		s.l.ErrPrintln(err)
		return nil, nil, err
	}

	return hashedPass, salt, nil
}

func (s Service) generateSalt() ([]byte, error) {
	var maxLength = new(big.Int)
	var maxValue = new(big.Int)
	var base = new(big.Int)

	maxLength.SetInt64(20)
	maxValue.SetInt64(93) // check utf-8 char set to understand this value
	base.SetInt64(32)     // check utf-8 char set to understand this value

	length, err := rand.Int(rand.Reader, maxLength)
	if err != nil {
		return nil, err
	}
	var salt = make([]byte, 0)

	for i := 0; i < int(length.Int64()); i++ {
		n, err := rand.Int(rand.Reader, maxValue)
		if err != nil {
			return nil, err
		}
		n.Add(n, base)
		salt = append(salt, byte(n.Int64()))
	}

	if len(salt) == 0 {
		return s.generateSalt()
	}

	return salt, nil
}

// maybe optimize this
// add tests for this!!!
// formula used: random(max - min) + min
func (s Service) GenCode() (uint64, error) {
	max := new(big.Int)
	max.SetInt64(999999)
	min := new(big.Int)
	min.SetInt64(100000)

	code, err := rand.Int(rand.Reader, max.Sub(max, min))
	if err != nil {
		return 0, err
	}

	code.Add(code, min)

	return code.Uint64(), nil
}

func (s Service) EncryptCode(code uint64) ([]byte, error) {
	c := []byte(strconv.Itoa(int(code)))
	hashedCode, err := bcrypt.GenerateFromPassword(c, bcrypt.DefaultCost)
	if err != nil {
		s.l.ErrPrintln(err)
		return nil, err
	}

	return hashedCode, err
}

// return true if code is corect
func (s Service) VerifyCode(realCode string, code uint64) (bool, error) {
	codeStr := []byte(strconv.Itoa(int(code)))
	err := bcrypt.CompareHashAndPassword([]byte(realCode), codeStr)
	if err != nil {
		return false, err
	}

	return true, nil
}
