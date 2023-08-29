package jwt

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	_ = os.Setenv("JWT_SECRET", "verysecret")

	id := int64(123)
	username := "testuser"
	email := "test@example.com"

	jwtImpl := JWTImpl{}

	tokenString, err := jwtImpl.Generate(id, username, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	assert.NoError(t, err)
	claims, ok := token.Claims.(*JWTClaim)
	assert.True(t, ok)
	assert.Equal(t, id, claims.ID)
	assert.Equal(t, username, claims.Username)
	assert.Equal(t, email, claims.Email)
	assert.True(t, claims.ExpiresAt.Unix() > time.Now().Unix())
}
