package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Ping(ctx context.Context) (err error) {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockRepo) MongoPing(ctx context.Context) (err error) {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockRepo) RedisPing(ctx context.Context) (err error) {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestHealthCheck(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("Ping", mock.Anything).Return(nil)
	mockRepo.On("MongoPing", mock.Anything).Return(nil)
	mockRepo.On("RedisPing", mock.Anything).Return(nil)

	u := NewUsecase(mockRepo)

	err := u.HealthCheck(context.Background())
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Ping", mock.Anything)
	mockRepo.AssertCalled(t, "MongoPing", mock.Anything)
	mockRepo.AssertCalled(t, "RedisPing", mock.Anything)
}

func TestHealthCheckErrorPing(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("Ping", mock.Anything).Return(errors.New("error ping"))
	mockRepo.On("MongoPing", mock.Anything).Return(nil)
	mockRepo.On("RedisPing", mock.Anything).Return(nil)

	u := NewUsecase(mockRepo)

	err := u.HealthCheck(context.Background())
	assert.Error(t, err)
}

func TestHealthCheckErrorMongoPing(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("Ping", mock.Anything).Return(nil)
	mockRepo.On("MongoPing", mock.Anything).Return(errors.New("error mongo ping"))
	mockRepo.On("RedisPing", mock.Anything).Return(nil)

	u := NewUsecase(mockRepo)

	err := u.HealthCheck(context.Background())
	assert.Error(t, err)
}

func TestHealthCheckErrorRedisPing(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("Ping", mock.Anything).Return(nil)
	mockRepo.On("MongoPing", mock.Anything).Return(nil)
	mockRepo.On("RedisPing", mock.Anything).Return(errors.New("error redis ping"))

	u := NewUsecase(mockRepo)

	err := u.HealthCheck(context.Background())
	assert.Error(t, err)
}
