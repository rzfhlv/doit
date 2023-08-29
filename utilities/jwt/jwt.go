package jwt

type JWTInterface interface {
	Generate(id int64, username, email string) (tokenString string, err error)
	ValidateToken(signedToken string) (claims *JWTClaim, err error)
}

type JWTImpl struct{}
