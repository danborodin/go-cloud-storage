package handlers

import (
	"auth/src/services/email"
	"auth/src/services/user"
	"auth/src/types"
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/danborodin/go-logd"
)

type RegisterHandler struct {
	l            *logd.Logger
	userService  *user.Service
	emailService email.Service
}

func NewRegHandler(l *logd.Logger, usrSrvc *user.Service, emailSrvc email.Service) *RegisterHandler {
	return &RegisterHandler{
		l:            l,
		userService:  usrSrvc,
		emailService: emailSrvc,
	}
}

func (h RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h RegisterHandler) post(w http.ResponseWriter, r *http.Request) error {
	var user = new(types.User)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)
	if err != nil {
		return err
	}

	err = h.userService.RegisterUser(user)
	if err != nil {
		return err
	}

	t, err := template.New("confirmationCode.html").ParseFiles("src/services/email/templates/confirmationCode.html")
	if err != nil {
		return err
	}
	data := struct {
		Username string
		Code     string
	}{
		Username: user.Username,
		Code:     strconv.Itoa(int(user.Verification.Code)),
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return err
	}

	to := []string{user.Email}
	subject := "Registration code"

	err = h.emailService.SendEmail(subject, buf.String(), to, "text/html")
	if err != nil {
		return err
	}

	return nil
}
