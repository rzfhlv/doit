package seeders

import (
	"testing"

	"github.com/rzfhlv/doit/config"
	"github.com/stretchr/testify/assert"
)

func TestNewSeed(t *testing.T) {
	cfg := config.Config{
		Postgres: nil,
		Mongo:    nil,
		Redis:    nil,
	}
	s := NewSeed(&cfg)
	assert.NotNil(t, s)
}
