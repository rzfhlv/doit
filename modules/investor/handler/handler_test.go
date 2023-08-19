package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rzfhlv/doit/modules/investor/model"
	mockUsecase "github.com/rzfhlv/doit/utilities/mocks/investor/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name      string
	wantError error
	err       bool
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
			name: "Testcase #1: Positive", wantError: nil, err: false, id: "?page=1", code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, err: true, id: "?page=1", code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative Param", wantError: errFoo, err: true, id: "?page=dua", code: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("GetAll", mock.Anything, mock.Anything).Return([]model.Investor{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/"+tt.id, nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := h.GetAll(ctx)
			if !tt.err {
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
			name: "Testcase #1: Positive", wantError: nil, err: false, id: "1", code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, err: true, id: "1", code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative Wrong ID", wantError: errFoo, err: true, id: "s", code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #4: Negative No Rows", wantError: sql.ErrNoRows, err: true, id: "0", code: http.StatusNotFound,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("GetByID", mock.Anything, mock.Anything).Return(model.Investor{}, tt.wantError)

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
			if !tt.err {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.code, rec.Code)
			}
		})
	}
}

func TestGenerateHandler(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, err: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, err: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Generate", mock.Anything).Return(tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/generate", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := h.Generate(ctx)
			if !tt.err {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
			}
		})
	}
}

func TestMigrateHandler(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, err: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, err: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("MigrateInvestors", mock.Anything).Return(tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/generate", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := h.Migrate(ctx)
			if !tt.err {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
			}
		})
	}
}
