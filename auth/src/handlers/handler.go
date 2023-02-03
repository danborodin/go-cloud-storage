package handlers

import (
	"auth/src/services"
	"github.com/danborodin/go-logd"
)

const rootUrl = "/api/v1/"

type RouterHandlers struct {
	RegisterHandler *RegisterHandler
}

func NewRouterHandlers(l *logd.Logger, srvc *services.Services) *RouterHandlers {
	return &RouterHandlers{
		RegisterHandler: NewRegHandler(l, srvc.UserService),
	}
}
