package handlers

import (
	"auth/src/services/email"
	"auth/src/services/user"
	"encoding/json"
	"net/http"

	"github.com/danborodin/go-logd"
)

type VerifyHandler struct {
	l            *logd.Logger
	userService  *user.Service
	emailService email.Service
}

type verifyHttpReq struct {
	Username string `json:"username"`
	Code     uint64 `json:"code"`
}

func NewVerifyHandler(l *logd.Logger, usrSrvc *user.Service, emailSrvc email.Service) *VerifyHandler {
	return &VerifyHandler{
		l:            l,
		userService:  usrSrvc,
		emailService: emailSrvc,
	}
}

func (h VerifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := h.post(w, r)
	if err != nil {
		h.l.ErrPrintln(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h VerifyHandler) post(w http.ResponseWriter, r *http.Request) error {
	var userVerification = new(verifyHttpReq)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(userVerification)
	if err != nil {
		return err
	}

	err = h.userService.VerifyUser(userVerification.Username, userVerification.Code)
	if err != nil {
		return err
	}

	return nil
}
