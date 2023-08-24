package investor

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rzfhlv/doit/config"
	mockAuth "github.com/rzfhlv/doit/utilities/mocks/auth"
	mockHandler "github.com/rzfhlv/doit/utilities/mocks/investor/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, target string
	wantError    error
	code         int
}

func TestMount(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", target: "/api/investors", wantError: nil, code: http.StatusOK,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := mockHandler.IHandler{}
			mockHandler.On("GetAll", mock.Anything).Return(tt.wantError)

			mockAuth := mockAuth.IAuth{}
			mockAuth.On("Bearer", mock.Anything).Return(func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					return next(c)
				}
			})

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			route := e.Group("/api")
			Mount(route, &mockHandler, &mockAuth)

			e.ServeHTTP(c.Response(), c.Request())

			assert.Equal(t, tt.code, rec.Code)
		})
	}
}

func TestNewInvestor(t *testing.T) {
	cfg := config.Config{
		Postgres: nil,
		Mongo:    nil,
		Redis:    nil,
	}

	investor := NewInvestor(&cfg)
	assert.NotNil(t, investor)
}
