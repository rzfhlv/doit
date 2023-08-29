package auth

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/labstack/echo/v4"
	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/utilities/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	jwtImpl := jwt.JWTImpl{}
	token, _ := jwtImpl.Generate(int64(1), "test", "test@example.com")
	client, mock := redismock.NewClientMock()
	mock.ExpectGet(token).SetVal("test")

	cfg := config.Config{
		Redis:   client,
		JWTImpl: &jwtImpl,
	}

	e := echo.New()
	auth := NewAuth(&cfg)
	e.Use(auth.Bearer)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}
	err := auth.Bearer(handler)(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test", rec.Body.String())

	assert.Equal(t, int64(1), c.Get("id"))
	assert.Equal(t, "test", c.Get("username"))
	assert.Equal(t, "test@example.com", c.Get("email"))
}

func TestFailAuthMiddleware(t *testing.T) {
	jwtImpl := jwt.JWTImpl{}
	token := "invalid-token"
	client, mock := redismock.NewClientMock()
	mock.ExpectGet(token).SetErr(errors.New("error"))

	cfg := config.Config{
		Redis:   client,
		JWTImpl: &jwtImpl,
	}

	e := echo.New()
	auth := NewAuth(&cfg)
	e.Use(auth.Bearer)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}
	auth.Bearer(handler)(c)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
