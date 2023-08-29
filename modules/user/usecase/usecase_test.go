package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rzfhlv/doit/modules/user/model"
	mockRepo "github.com/rzfhlv/doit/shared/mocks/modules/user/repository"
	mockHasher "github.com/rzfhlv/doit/shared/mocks/utilities/hasher"
	mockJwt "github.com/rzfhlv/doit/shared/mocks/utilities/jwt"
	"github.com/rzfhlv/doit/utilities/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name                                                                string
	wantError, wantErrorUsername, wantErrJwt, wantErrHash, wantErrRedis error
	token                                                               model.Validate
	isErr                                                               bool
}

var (
	errFoo      = errors.New("error")
	tokenString string
	claims      *jwt.JWTClaim

	testUser = model.User{
		ID:        1,
		Name:      "testuser",
		Email:     "test@example.com",
		Username:  "testuser",
		Password:  "verysecret",
		CreatedAt: time.Now(),
	}

	testLogin = model.Login{
		Username: "testuser",
		Password: "verysecret",
	}
)

func TestNewUsecase(t *testing.T) {
	mockRepo := mockRepo.IRepository{}
	mockJwt := mockJwt.JWTInterface{}
	mockHasher := mockHasher.HashPassword{}

	u := NewUsecase(&mockRepo, &mockJwt, &mockHasher)
	assert.NotNil(t, u)
}

func TestRegisterUsecase(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, wantErrorUsername: errFoo, wantErrJwt: nil, wantErrHash: nil, wantErrRedis: nil, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantErrorUsername: errFoo, wantErrJwt: nil, wantErrHash: nil, wantErrRedis: nil, isErr: true,
		},
		{
			name: "Testcase #3: Negative", wantError: nil, wantErrorUsername: nil, wantErrJwt: nil, wantErrHash: nil, wantErrRedis: nil, isErr: true,
		},
		{
			name: "Testcase #4: Negative", wantError: nil, wantErrorUsername: errFoo, wantErrJwt: errFoo, wantErrHash: nil, wantErrRedis: nil, isErr: true,
		},
		{
			name: "Testcase #4: Negative", wantError: nil, wantErrorUsername: errFoo, wantErrJwt: nil, wantErrHash: errFoo, wantErrRedis: nil, isErr: true,
		},
		{
			name: "Testcase #4: Negative", wantError: nil, wantErrorUsername: errFoo, wantErrJwt: nil, wantErrHash: nil, wantErrRedis: errFoo, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("Login", mock.Anything, mock.Anything).Return(model.User{}, tt.wantErrorUsername)
			mockRepo.On("Register", mock.Anything, mock.Anything).Return(testUser, tt.wantError)
			mockRepo.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.wantErrRedis)

			mockJwt := mockJwt.JWTInterface{}
			mockJwt.On("Generate", mock.Anything, mock.Anything, mock.Anything).Return(tokenString, tt.wantErrJwt)

			mockHasher := mockHasher.HashPassword{}
			mockHasher.On("HashedPassword", mock.Anything).Return("", tt.wantErrHash)

			u := &Usecase{
				repo:    &mockRepo,
				jwtImpl: &mockJwt,
				hasher:  &mockHasher,
			}

			_, err := u.Register(context.Background(), testUser)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLoginUsecase(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, wantErrJwt: nil, wantErrHash: nil, wantErrRedis: nil, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantErrJwt: nil, wantErrHash: nil, wantErrRedis: nil, isErr: true,
		},
		{
			name: "Testcase #3: Negative", wantError: nil, wantErrJwt: errFoo, wantErrHash: nil, wantErrRedis: nil, isErr: true,
		},
		{
			name: "Testcase #4: Negative", wantError: nil, wantErrJwt: nil, wantErrHash: bcrypt.ErrMismatchedHashAndPassword, wantErrRedis: nil, isErr: true,
		},
		{
			name: "Testcase #5: Negative", wantError: nil, wantErrJwt: nil, wantErrHash: nil, wantErrRedis: errFoo, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("Login", mock.Anything, testLogin).Return(model.User{}, tt.wantError)
			mockRepo.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.wantErrRedis)

			mockJwt := mockJwt.JWTInterface{}
			mockJwt.On("Generate", mock.Anything, mock.Anything, mock.Anything).Return(tokenString, tt.wantErrJwt)

			mockHasher := mockHasher.HashPassword{}
			mockHasher.On("VerifyPassword", mock.Anything, mock.Anything).Return(tt.wantErrHash)

			u := &Usecase{
				repo:    &mockRepo,
				jwtImpl: &mockJwt,
				hasher:  &mockHasher,
			}

			_, err := u.Login(context.Background(), testLogin)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestValidateUsecase(t *testing.T) {
	jwtImpl := jwt.JWTImpl{}
	token, _ := jwtImpl.Generate(int64(1), "test", "test@example.com")
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, token: model.Validate{Token: token}, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, token: model.Validate{Token: "invalid"}, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockJwt := mockJwt.JWTInterface{}
			mockJwt.On("ValidateToken", mock.Anything).Return(claims, tt.wantError)

			mockHaser := mockHasher.HashPassword{}

			u := &Usecase{
				repo:    &mockRepo,
				jwtImpl: &mockJwt,
				hasher:  &mockHaser,
			}

			_, err := u.Validate(context.Background(), tt.token)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLogoutUsecase(t *testing.T) {
	token := "thisistoken"
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("Del", mock.Anything, mock.Anything).Return(tt.wantError)

			mockJwt := mockJwt.JWTInterface{}
			mockHasher := mockHasher.HashPassword{}

			u := &Usecase{
				repo:    &mockRepo,
				jwtImpl: &mockJwt,
				hasher:  &mockHasher,
			}

			err := u.Logout(context.Background(), token)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
