package auth

import (
	"context"
	"doit/utilities"
	"doit/utilities/jwt"
	"fmt"
	"net/http"
	"strings"

	logrus "doit/utilities/log"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
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
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		if split[0] != "Bearer" {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Unsupported Token: %v", split[0]))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		if split[1] == "" {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Empty Token: %v", split[1]))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		claims, err := jwt.ValidateToken(split[1])
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Validation Invalid, %v", err.Error()))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		err = am.redis.Get(context.Background(), split[1]).Err()
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf("Auth Redis Key Deleted, %v", err.Error()))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)

		return next(c)
	}
}
