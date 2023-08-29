package middleware

import (
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/utilities/jwt"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	jwtImpl := jwt.JWTImpl{}
	client, _ := redismock.NewClientMock()

	cfg := config.Config{
		Redis:   client,
		JWTImpl: &jwtImpl,
	}

	m := NewMiddleware(&cfg)
	assert.NotNil(t, m)
}
