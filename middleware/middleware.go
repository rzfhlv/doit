package middleware

import (
	"doit/middleware/auth"

	"github.com/redis/go-redis/v9"
)

type Middleware struct {
	Auth auth.IAuth
}

func NewMiddleware(redis *redis.Client) *Middleware {
	auth := auth.NewAuth(redis)

	return &Middleware{
		Auth: auth,
	}
}
