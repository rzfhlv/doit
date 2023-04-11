package auth

import (
	"context"
	"doit/utilities"
	"doit/utilities/jwt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type IAuthMiddleware interface {
	AuthBearer(next echo.HandlerFunc) echo.HandlerFunc
}

type AuthMiddleware struct {
	redis *redis.Client
}

func NewAuthMiddleware(redis *redis.Client) IAuthMiddleware {
	return &AuthMiddleware{
		redis: redis,
	}
}

func (am *AuthMiddleware) AuthBearer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		split := strings.Split(c.Request().Header.Get("Authorization"), " ")
		if len(split) < 2 {
			log.Printf("[ERROR] Auth Unsupported Token: %v", len(split))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		if split[0] != "Bearer" {
			log.Printf("[ERROR] Auth Unsupported Token: %v", split[0])
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		if split[1] == "" {
			log.Printf("[ERROR] Auth Empty Token: %v", split[1])
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		claims, err := jwt.ValidateToken(split[1])
		if err != nil {
			log.Printf("[ERROR] Auth Validation Invalid: %v", err.Error())
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		err = am.redis.Get(context.Background(), split[1]).Err()
		if err != nil {
			log.Printf("[ERROR] Auth Redis Key Deleted: %v", err.Error())
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)

		return next(c)
	}
}
