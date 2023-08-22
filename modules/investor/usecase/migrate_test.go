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

// type MockUsecase struct {
// 	mock.Mock
// }

// func (m *MockUsecase) getInvestors() <-chan model.Investor {
// 	args := m.Called()
// 	return args.Get(0).(<-chan model.Investor)
// }

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

func TestMigrateInvestorsUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mockRepo := mockRepo.IRepository{}
		mockRepo.On("GetPsql", mock.Anything).Return([]model.Investor{}, nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		err := u.MigrateInvestors(context.Background())
		assert.NoError(t, err)
	})
}

func TestGetInvestorsUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mockRepo := mockRepo.IRepository{}
		mockInvestors := []model.Investor{
			{
				ID: int64(1), Name: "Test 1",
			},
			{
				ID: int64(2), Name: "Test 2",
			},
		}
		mockRepo.On("GetPsql", mock.Anything).Return(mockInvestors, nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		resultChan := u.getInvestors()

		var collectedInvestors []model.Investor
		for investor := range resultChan {
			collectedInvestors = append(collectedInvestors, investor)
		}

		assert.Equal(t, mockInvestors, collectedInvestors)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpsertInvestorsUsecase(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mockRepo := mockRepo.IRepository{}
		mockInvestor := model.Investor{
			ID: int64(1), Name: "Test 1",
		}
		mockRepo.On("UpsertOutbox", mock.Anything, mock.Anything).Return(nil)
		mockRepo.On("UpsertMongo", mock.Anything, mock.Anything).Return(nil)
		mockRepo.On("DeleteOutbox", mock.Anything, mock.Anything).Return(nil)

		u := &Usecase{
			repo: &mockRepo,
		}

		investorChan := make(chan model.Investor, 1)
		investorChan <- mockInvestor
		close(investorChan)

		resultChan := u.upsertInvestors(investorChan)

		var collectedInvestors []model.Investor
		for investor := range resultChan {
			collectedInvestors = append(collectedInvestors, investor)
		}

		assert.Len(t, collectedInvestors, 1)
	})
}

func TestMergeChanInvestorUsecase(t *testing.T) {
	mockRepo := mockRepo.IRepository{}

	mockData1 := []model.Investor{
		{
			ID: int64(1), Name: "Test 1",
		},
		{
			ID: int64(2), Name: "Test 2",
		},
	}
	mockData2 := []model.Investor{
		{
			ID: int64(3), Name: "Test 3",
		},
		{
			ID: int64(4), Name: "Test 4",
		},
	}

	chan1 := make(chan model.Investor)
	chan2 := make(chan model.Investor)

	go func() {
		defer close(chan1)
		for _, data := range mockData1 {
			chan1 <- data
		}
	}()
	go func() {
		defer close(chan2)
		for _, data := range mockData2 {
			chan2 <- data
		}
	}()

	u := &Usecase{
		repo: &mockRepo,
	}

	resultChan := u.mergeChanInvestor(chan1, chan2)
	var collectedData []model.Investor
	for data := range resultChan {
		collectedData = append(collectedData, data)
	}

	assert.ElementsMatch(t, append(mockData1, mockData2...), collectedData)
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
