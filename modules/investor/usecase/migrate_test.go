package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rzfhlv/doit/modules/investor/model"
	mockRepo "github.com/rzfhlv/doit/utilities/mocks/investor/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestConventionalMigrateUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("GetPsql", mock.Anything).Return([]model.Investor{}, nil)
		mockRepo.On("UpsertMongo", mock.Anything, mock.Anything).Return(nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.ConventionalMigrate(context.Background())
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative GetPsql", func(t *testing.T) {
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("GetPsql", mock.Anything).Return([]model.Investor{}, errors.New("error"))
		mockRepo.On("UpsertMongo", mock.Anything, mock.Anything).Return(nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.ConventionalMigrate(context.Background())
		assert.Error(t, err)
	})
}

func TestSaveOutBoxUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		var payload []byte
		now := time.Now()
		investor := model.Investor{}

		mockRepo := mockRepo.IRepository{}
		mockRepo.On("UpsertOutbox", mock.Anything, mock.Anything).Return(nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.saveOutbox(context.Background(), now, payload, investor)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		var payload []byte
		now := time.Now()
		investor := model.Investor{}

		mockRepo := mockRepo.IRepository{}
		mockRepo.On("UpsertOutbox", mock.Anything, mock.Anything).Return(errors.New("error"))

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.saveOutbox(context.Background(), now, payload, investor)
		assert.Error(t, err)
	})
}

func TestUpsertMongoUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		var payload []byte
		now := time.Now()
		investor := model.Investor{}

		mockRepo := mockRepo.IRepository{}
		mockRepo.On("UpsertMongo", mock.Anything, mock.Anything).Return(nil)
		mockRepo.On("UpsertOutbox", mock.Anything, mock.Anything).Return(nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.upsertMongo(context.Background(), now, payload, investor)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative UpsertMongo", func(t *testing.T) {
		var payload []byte
		now := time.Now()
		investor := model.Investor{}

		mockRepo := mockRepo.IRepository{}
		mockRepo.On("UpsertMongo", mock.Anything, mock.Anything).Return(errors.New("error"))
		mockRepo.On("UpsertOutbox", mock.Anything, mock.Anything).Return(errors.New("error"))

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.upsertMongo(context.Background(), now, payload, investor)
		assert.Error(t, err)
	})
}

func TestDeleteOutboxUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		var id int64 = 1

		mockRepo := mockRepo.IRepository{}
		mockRepo.On("DeleteOutbox", mock.Anything, mock.Anything).Return(nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.deleteOutbox(context.Background(), id)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		var id int64 = 1

		mockRepo := mockRepo.IRepository{}
		mockRepo.On("DeleteOutbox", mock.Anything, mock.Anything).Return(errors.New("error"))

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.deleteOutbox(context.Background(), id)
		assert.Error(t, err)
	})
}
