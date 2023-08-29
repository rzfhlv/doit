package middleware

import (
	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/middleware/auth"
	"github.com/rzfhlv/doit/middleware/log"
)

type Middleware struct {
	Auth auth.IAuth
	Log  log.ILog
}

func NewMiddleware(cfg *config.Config) *Middleware {
	auth := auth.NewAuth(cfg)
	log := log.NewLog()

	return &Middleware{
		Auth: auth,
		Log:  log,
	}
}
