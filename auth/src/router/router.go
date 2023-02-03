package router

import (
	"auth/src/handlers"
	"auth/src/router/middleware"
	"github.com/danborodin/go-logd"
	"net/http"
)

var md *middleware.Middleware

func CreateRouter(l *logd.Logger, handlers *handlers.RouterHandlers) http.Handler {
	md = middleware.NewMd(l)

	muxTest := http.NewServeMux()
	muxTest.Handle("/api/v1/register", md.TestMd(handlers.RegisterHandler))
	muxTest.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("world hello?"))
	})

	return muxTest
}
