package middleware

import (
	"github.com/rzfhlv/doit/middleware/auth"
	"github.com/rzfhlv/doit/middleware/log"

	"github.com/redis/go-redis/v9"
)

type Middleware struct {
	Auth auth.IAuth
	Log  log.ILog
}

func NewMiddleware(redis *redis.Client) *Middleware {
	auth := auth.NewAuth(redis)
	log := log.NewLog()

	return &Middleware{
		Auth: auth,
		Log:  log,
	}
}
