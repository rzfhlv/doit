package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/rzfhlv/doit/modules/person/model"
	mockRepo "github.com/rzfhlv/doit/shared/mocks/modules/person/repository"
	"github.com/rzfhlv/doit/utilities/param"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name                      string
	wantError, wantCountError error
	isErr                     bool
	id                        int64
}

var (
	errFoo = errors.New("error")
)

func TestNewUsecase(t *testing.T) {
	mockRepo := mockRepo.IRepository{}

	u := NewUsecase(&mockRepo)
	assert.NotNil(t, u)
}

func TestGetAllUsecase(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, wantCountError: nil, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantCountError: nil, isErr: true,
		},
		{
			name: "Testcase #3: Negative", wantError: nil, wantCountError: errFoo, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("GetAll", mock.Anything, mock.Anything).Return([]model.Person{}, tt.wantError)
			mockRepo.On("Count", mock.Anything).Return(int64(10), tt.wantCountError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.GetAll(context.Background(), &param.Param{})
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetByIDUsecase(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, isErr: false, id: 1,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, isErr: true, id: 1,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(model.Person{}, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.GetByID(context.Background(), tt.id)
			if !tt.isErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
