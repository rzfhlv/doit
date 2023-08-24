package healthcheck

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rzfhlv/doit/config"
	mockHandler "github.com/rzfhlv/doit/utilities/mocks/health-check/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, target string
	wantError    error
	code         int
}

var (
	errFoo = errors.New("error")
)

func TestMount(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", target: "/api/health-check", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", target: "/api/health-check", wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #2: Negative", target: "/api/health", wantError: errFoo, code: http.StatusNotFound,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := mockHandler.IHandler{}
			mockHandler.On("HealthCheck", mock.Anything).Return(tt.wantError)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			route := e.Group("/api")
			Mount(route, &mockHandler)

			e.ServeHTTP(c.Response(), c.Request())

			assert.Equal(t, tt.code, rec.Code)
		})
	}
}

func TestNewHealthCheck(t *testing.T) {
	cfg := config.Config{
		Postgres: nil,
		Mongo:    nil,
		Redis:    nil,
	}

	hc := NewHealthCheck(&cfg)
	assert.NotNil(t, hc)
}
