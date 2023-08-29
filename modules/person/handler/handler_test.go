package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rzfhlv/doit/modules/person/model"
	mockUsecase "github.com/rzfhlv/doit/shared/mocks/modules/person/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type testCase struct {
	name      string
	wantError error
	isErr     bool
	id        string
	code      int
}

var (
	errFoo = errors.New("error")
)

func TestNewHandler(t *testing.T) {
	mockUsecase := mockUsecase.IUsecase{}

	h := NewHandler(&mockUsecase)
	assert.NotNil(t, h)
}

func TestGetAllHandler(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, isErr: false, id: "?page=1", code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, isErr: true, id: "?page=1", code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", wantError: errFoo, isErr: true, id: "?page=satu", code: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("GetAll", mock.Anything, mock.Anything).Return([]model.Person{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/"+tt.id, nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := h.GetAll(ctx)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.code, rec.Code)
			}
		})
	}
}

func TestGetByIDHandler(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, isErr: false, id: "1", code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, isErr: true, id: "1", code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", wantError: errFoo, isErr: true, id: "s", code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #4: Negative", wantError: mongo.ErrNoDocuments, isErr: true, id: "0", code: http.StatusNotFound,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("GetByID", mock.Anything, mock.Anything).Return(model.Person{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/"+tt.id, nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.id)

			err := h.GetByID(ctx)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.code, rec.Code)
			}
		})
	}
}
