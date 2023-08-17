package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) HealthCheck(ctx context.Context) (err error) {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestHealthCheck(t *testing.T) {
	mockUsecase := new(MockUsecase)
	mockUsecase.On("HealthCheck", mock.Anything).Return(nil)

	h := NewHandler(mockUsecase)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := h.HealthCheck(ctx)
	assert.NoError(t, err)
	mockUsecase.AssertCalled(t, "HealthCheck", mock.Anything)
}

func TestHealthCheckError(t *testing.T) {
	mockUsecase := new(MockUsecase)
	mockUsecase.On("HealthCheck", mock.Anything).Return(errors.New("error usecase"))

	h := NewHandler(mockUsecase)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	_ = h.HealthCheck(ctx)
	mockUsecase.AssertExpectations(t)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
