package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var (
	now   = time.Date(2023, time.August, 15, 12, 0, 0, 0, time.UTC)
	ctx   = context.Background()
	ttl   = time.Duration(1 * time.Hour)
	key   = "testKey"
	value = "testValue"
)

func TestNewRepository(t *testing.T) {
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	client, _ := redismock.NewClientMock()

	r := NewRepository(db, client)
	assert.NotNil(t, r)
}
