package handlers

import (
	"auth/src/services/user"
	"auth/src/types"
	"encoding/json"
	"net/http"

	"github.com/danborodin/go-logd"
)

type RegisterHandler struct {
	l           *logd.Logger
	userService *user.Service
}

func NewRegHandler(l *logd.Logger, usrSrvc *user.Service) *RegisterHandler {
	return &RegisterHandler{
		l:           l,
		userService: usrSrvc,
	}
}

func (rh RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user = new(types.User)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)
	if err != nil {
		rh.l.ErrPrintln(err)
		w.WriteHeader(http.StatusBadRequest)
		//w.Write() // add a detailed error
		return
	}

	err = rh.userService.RegisterUser(user)
	if err != nil {
		rh.l.ErrPrintln(err)
	}
}
