package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok && !token.Valid {
		err = errors.New("couldn't parse claims")
		return
	}
	return
}
