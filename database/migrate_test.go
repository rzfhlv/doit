package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/doit/config"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	mockDB, _, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	cfg := config.Config{
		Postgres: db,
		Redis:    nil,
		Mongo:    nil,
	}
	Migrate(&cfg)
	assert.True(t, true)
}
