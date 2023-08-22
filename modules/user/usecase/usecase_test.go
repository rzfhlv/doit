package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rzfhlv/doit/modules/user/model"
	mockRepo "github.com/rzfhlv/doit/utilities/mocks/user/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name                         string
	wantError, wantErrorUsername error
	isErr                        bool
}

var (
	errFoo = errors.New("error")
)

func TestNewUsecase(t *testing.T) {
	mockRepo := mockRepo.IRepository{}

	u := NewUsecase(&mockRepo)
	assert.NotNil(t, u)
}

func TestRegisterUsecase(t *testing.T) {
	testUser := model.User{
		ID:        1,
		Name:      "testuser",
		Email:     "test@example.com",
		Username:  "testuser",
		Password:  "verysecret",
		CreatedAt: time.Now(),
	}
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, wantErrorUsername: errFoo, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantErrorUsername: errFoo, isErr: true,
		},
		{
			name: "Testcase #3: Negative", wantError: nil, wantErrorUsername: nil, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("Login", mock.Anything, mock.Anything).Return(model.User{}, tt.wantErrorUsername)
			mockRepo.On("Register", mock.Anything, mock.Anything).Return(testUser, tt.wantError)
			mockRepo.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
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
	testLogin := model.Login{
		Username: "testuser",
		Password: "verysecret",
	}
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
			mockRepo.On("Login", mock.Anything, mock.Anything).Return(model.User{}, tt.wantError)
			mockRepo.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
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

			u := &Usecase{
				repo: &mockRepo,
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
