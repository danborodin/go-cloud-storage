package middleware

import (
	"net/http"

	"github.com/danborodin/go-logd"
)

type Middleware struct {
	l *logd.Logger
}

func NewMd(l *logd.Logger) *Middleware {
	return &Middleware{
		l: l,
	}
}

func (md *Middleware) TestMd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md.l.InfoPrintln("create user req")
		next.ServeHTTP(w, r)
		return
	})
}
