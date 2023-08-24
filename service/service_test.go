package service

import (
	"testing"

	"github.com/rzfhlv/doit/config"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	cfg := config.Config{
		Postgres: nil,
		Mongo:    nil,
		Redis:    nil,
	}

	s := NewService(&cfg)
	assert.NotNil(t, s)
}
