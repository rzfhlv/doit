package middleware

import (
	"doit/middleware/auth"
	"doit/middleware/log"

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
