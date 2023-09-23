package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rzfhlv/doit/config"
	uJwt "github.com/rzfhlv/doit/utilities/jwt"
	"github.com/rzfhlv/doit/utilities/message"
	"github.com/rzfhlv/doit/utilities/response"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

const (
	ID       = "id"
	EMAIL    = "email"
	USERNAME = "username"

	BEARER        = "Bearer"
	AUTHORIZATION = "Authorization"

	UNSUPPORTEDTOKENLOG  = "Auth Unsupported Token"
	EMPTYTOKENLOG        = "Auth Empty Token"
	VALIDATIONINVALIDLOG = "Auth Validation Invalid"
	REDISLOG             = "Auth Redis Key Deleted"
)

type IAuth interface {
	Bearer(next echo.HandlerFunc) echo.HandlerFunc
}

type Auth struct {
	redis   *redis.Client
	jwtImpl uJwt.JWTInterface
}

func NewAuth(cfg *config.Config) IAuth {
	return &Auth{
		redis:   cfg.Redis,
		jwtImpl: cfg.JWTImpl,
	}
}

func (am *Auth) Bearer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		split := strings.Split(c.Request().Header.Get(AUTHORIZATION), " ")
		if len(split) < 2 {
			logrus.Log(nil).Error(fmt.Sprintf(UNSUPPORTEDTOKENLOG+" %v", len(split)))
			return c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		if split[0] != BEARER {
			logrus.Log(nil).Error(fmt.Sprintf(UNSUPPORTEDTOKENLOG+" %v", split[0]))
			return c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		if split[1] == "" {
			logrus.Log(nil).Error(fmt.Sprintf(EMPTYTOKENLOG+" %v", split[1]))
			return c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		claims, err := am.jwtImpl.ValidateToken(split[1])
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf(VALIDATIONINVALIDLOG+" %v", err.Error()))
			return c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		err = am.redis.Get(context.Background(), split[1]).Err()
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf(REDISLOG+" %v", err.Error()))
			return c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		c.Set(ID, claims.ID)
		c.Set(EMAIL, claims.Email)
		c.Set(USERNAME, claims.Username)

		return next(c)
	}
}
