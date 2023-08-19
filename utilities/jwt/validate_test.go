package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateToken_Valid(t *testing.T) {
	validToken, _ := Generate(123, "testuser", "test@example.com")

	claims, err := ValidateToken(validToken)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, int64(123), claims.ID)
	assert.Equal(t, "testuser", claims.Username)
	assert.Equal(t, "test@example.com", claims.Email)
}

func TestValidateToken_Invalid(t *testing.T) {
	invalidToken := "invalidtoken"

	claims, err := ValidateToken(invalidToken)
	assert.Error(t, err)
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, "token is malformed: token contains an invalid number of segments", err.Error())
}
