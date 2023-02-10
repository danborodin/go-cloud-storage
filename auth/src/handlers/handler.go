package handlers

import (
	"auth/src/services"

	"github.com/danborodin/go-logd"
)

const rootUrl = "/api/v1/"

type Handlers struct {
	RegisterHandler *RegisterHandler
	VerifyHandler   *VerifyHandler
}

func NewHandlers(l *logd.Logger, srvc *services.Services) *Handlers {
	return &Handlers{
		RegisterHandler: NewRegHandler(l, srvc.UserService, srvc.EmailService),
		VerifyHandler:   NewVerifyHandler(l, srvc.UserService, srvc.EmailService),
	}
}
