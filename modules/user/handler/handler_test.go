package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rzfhlv/doit/modules/user/model"
	"github.com/rzfhlv/doit/utilities/jwt"
	"github.com/rzfhlv/doit/utilities/message"
	mockUsecase "github.com/rzfhlv/doit/utilities/mocks/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, reqBody string
	wantError     error
	isErr         bool
	code          int
}

var (
	errFoo = errors.New("error")
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func TestNewHandler(t *testing.T) {
	mockUsecase := mockUsecase.IUsecase{}

	h := NewHandler(&mockUsecase)
	assert.NotNil(t, h)
}

func TestRegisterHandler(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: `{"name": "testuser", "email": "test@example.com", "username": "testuser", "password": "secret"}`, wantError: nil, isErr: false, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqBody: `{"name": "testuser", "email": "test@example.com", "username": "testuser", "password": "secret"}`, wantError: errFoo, isErr: true, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", reqBody: `{"name": 123, "email": "test@example.com", "username": "testuser", "password": "secret"}`, wantError: errFoo, isErr: true, code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #4: Negative", reqBody: `{"name": "", "email": "test@example.com", "username": "testuser", "password": "secret"}`, wantError: errFoo, isErr: true, code: http.StatusBadRequest,
		},
		{
			name: "Testcase #5: Negative", reqBody: `{"name": "test", "email": "test@example.com", "username": "testuser", "password": "secret"}`, wantError: errors.New(message.ERRUSERNAMEEXIST), isErr: true, code: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Register", mock.Anything, mock.Anything).Return(model.JWT{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			t.Log(tt.wantError)

			err := h.Register(ctx)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.code, rec.Code)
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: `{"username": "testuser", "password": "password"}`, wantError: nil, isErr: false, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqBody: `{"username": "testuser", "password": "password"}`, wantError: errFoo, isErr: true, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", reqBody: `{"username": 123, "password": "password"}`, wantError: errFoo, isErr: true, code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #4: Negative", reqBody: `{"username": "", "password": "password"}`, wantError: errFoo, isErr: true, code: http.StatusBadRequest,
		},
		{
			name: "Testcase #5: Negative", reqBody: `{"username": "testuser", "password": "password"}`, wantError: sql.ErrNoRows, isErr: true, code: http.StatusBadRequest,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Login", mock.Anything, mock.Anything).Return(model.JWT{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := h.Login(ctx)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.code, rec.Code)
			}
		})
	}
}

func TestValidateUsecase(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: `{"token": "thisistoken"}`, wantError: nil, isErr: false, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqBody: `{"token": "thisistoken"}`, wantError: errFoo, isErr: true, code: http.StatusUnauthorized,
		},
		{
			name: "Testcase #3: Negative", reqBody: `{"token": 123}`, wantError: errFoo, isErr: true, code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #3: Negative", reqBody: `{"token": ""}`, wantError: errFoo, isErr: true, code: http.StatusBadRequest,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := jwt.JWTClaim{}
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Validate", mock.Anything, mock.Anything).Return(&result, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := h.Validate(ctx)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.code, rec.Code)
			}
		})
	}
}

func TestLogoutHandler(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: "Bearer thisistoken", wantError: nil, isErr: false, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqBody: "Bearer thisistoken", wantError: errFoo, isErr: true, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #2: Negative", reqBody: "", wantError: errFoo, isErr: true, code: http.StatusUnauthorized,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Logout", mock.Anything, mock.Anything).Return(tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(echo.HeaderAuthorization, tt.reqBody)
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := h.Logout(ctx)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.code, rec.Code)
			}
		})
	}
}
