package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rzfhlv/doit/utilities"
	"github.com/rzfhlv/doit/utilities/jwt"
	"github.com/rzfhlv/doit/utilities/message"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

const (
	ID       = "id"
	EMAIL    = "email"
	USERNAME = "username"

	BEARER = "Bearer"
)

type IAuth interface {
	Bearer(next echo.HandlerFunc) echo.HandlerFunc
}

type Auth struct {
	redis *redis.Client
}

func NewAuth(redis *redis.Client) IAuth {
	return &Auth{
		redis: redis,
	}
}

func (am *Auth) Bearer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		split := strings.Split(c.Request().Header.Get("Authorization"), " ")
		if len(split) < 2 {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Unsupported Token: %v", len(split)))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		if split[0] != BEARER {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Unsupported Token: %v", split[0]))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		if split[1] == "" {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Empty Token: %v", split[1]))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		claims, err := jwt.ValidateToken(split[1])
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Validation Invalid, %v", err.Error()))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		err = am.redis.Get(context.Background(), split[1]).Err()
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Redis Key Deleted, %v", err.Error()))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}

		c.Set(ID, claims.ID)
		c.Set(EMAIL, claims.Email)
		c.Set(USERNAME, claims.Username)

		return next(c)
	}
}
