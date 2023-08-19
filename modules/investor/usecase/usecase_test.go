package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/rzfhlv/doit/modules/investor/model"
	mockRepo "github.com/rzfhlv/doit/utilities/mocks/investor/repository"
	"github.com/rzfhlv/doit/utilities/param"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUsecase(t *testing.T) {
	mockRepo := mockRepo.IRepository{}
	u := NewUsecase(&mockRepo)
	assert.NotNil(t, u)
}

func TestGetAllUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive GetAll", func(t *testing.T) {
		var count int64 = 10
		investors := []model.Investor{}
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("GetAll", mock.Anything, mock.Anything).Return([]model.Investor{}, nil)
		mockRepo.On("Count", mock.Anything).Return(count, nil)

		u := &Usecase{
			repo: &mockRepo,
		}
		result, err := u.GetAll(context.Background(), &param.Param{})
		assert.Equal(t, investors, result)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative GetAll", func(t *testing.T) {
		var count int64 = 10
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("GetAll", mock.Anything, mock.Anything).Return([]model.Investor{}, errors.New("error"))
		mockRepo.On("Count", mock.Anything).Return(count, nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		_, err := u.GetAll(context.Background(), &param.Param{})
		assert.Error(t, err)
	})

	t.Run("Testcase #3: Negative Count", func(t *testing.T) {
		var count int64 = 10
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("GetAll", mock.Anything, mock.Anything).Return([]model.Investor{}, nil)
		mockRepo.On("Count", mock.Anything).Return(count, errors.New("error"))

		u := &Usecase{
			repo: &mockRepo,
		}

		_, err := u.GetAll(context.Background(), &param.Param{})
		assert.Error(t, err)
	})
}

func TestGetByIDUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		var id int64 = 1
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(model.Investor{}, nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		_, err := u.GetByID(context.Background(), id)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		var id int64 = 1
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(model.Investor{}, errors.New("error"))

		u := &Usecase{
			repo: &mockRepo,
		}

		_, err := u.GetByID(context.Background(), id)
		assert.Error(t, err)
	})
}

func TestGenerateUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("Generate", mock.Anything, mock.Anything).Return(nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.Generate(context.Background())
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("Generate", mock.Anything, mock.Anything).Return(errors.New("error"))

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.Generate(context.Background())
		assert.Error(t, err)
	})
}
