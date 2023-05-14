package middleware

import (
	aMiddleware "doit/middleware/auth"

	"github.com/redis/go-redis/v9"
)

type Middleware struct {
	AuthMiddleware aMiddleware.IAuthMiddleware
}

func NewMiddleware(redis *redis.Client) *Middleware {
	authMiddleware := aMiddleware.NewAuthMiddleware(redis)

	return &Middleware{
		AuthMiddleware: authMiddleware,
	}
}
