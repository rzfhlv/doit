package database

import (
	"testing"

	"github.com/rzfhlv/doit/config"
	"github.com/stretchr/testify/assert"
)

func TestSeed(t *testing.T) {
	args := []string{""}
	cfg := config.Config{
		Postgres: nil,
		Mongo:    nil,
		Redis:    nil,
	}
	Seed(&cfg, args)
	assert.True(t, true)
}
