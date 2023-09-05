package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaim struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func (j *JWTImpl) Generate(id int64, username, email string) (tokenString string, err error) {
	exp, err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		return
	}
	expirationTime := time.Now().Add(time.Duration(exp) * time.Hour)
	claims := &JWTClaim{
		ID:       id,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("APP_NAME"),
			Subject:   username,
			ID:        strconv.FormatInt(id, 10),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
