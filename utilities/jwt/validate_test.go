package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateToken_Valid(t *testing.T) {
	jwtImpl := JWTImpl{}
	validToken, _ := jwtImpl.Generate(123, "testuser", "test@example.com")

	claims, err := jwtImpl.ValidateToken(validToken)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, int64(123), claims.ID)
	assert.Equal(t, "testuser", claims.Username)
	assert.Equal(t, "test@example.com", claims.Email)
}

func TestValidateToken_Invalid(t *testing.T) {
	invalidToken := "invalidtoken"

	jwtImpl := JWTImpl{}
	claims, err := jwtImpl.ValidateToken(invalidToken)
	assert.Error(t, err)
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, "token is malformed: token contains an invalid number of segments", err.Error())
}
