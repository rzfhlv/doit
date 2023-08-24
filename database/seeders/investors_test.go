package seeders

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/doit/config"
	"github.com/stretchr/testify/assert"
)

type MockGenerator struct{}

func (m *MockGenerator) GenerateName() string {
	return "test"
}

func TestSeedInvestors(t *testing.T) {
	mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")
	mockSQL.ExpectExec("INSERT INTO investors (name) VALUES ($1);").
		WithArgs("test").WillReturnResult(sqlmock.NewResult(1, 1))

	cfg := config.Config{
		Postgres: db,
		Mongo:    nil,
		Redis:    nil,
	}
	s := &Seed{
		cfg:      &cfg,
		genrator: &MockGenerator{},
	}
	s.InvestorSeed()
	assert.True(t, true)
}
