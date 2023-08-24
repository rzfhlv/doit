package middleware

import (
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	client, _ := redismock.NewClientMock()

	m := NewMiddleware(client)
	assert.NotNil(t, m)
}
