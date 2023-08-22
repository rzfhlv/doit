package repository

import (
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestSetRepository(t *testing.T) {
	client, mock := redismock.NewClientMock()
	r := &Repository{
		redis: client,
	}

	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock.ExpectSet(key, value, ttl).SetVal("OK")
		err := r.Set(ctx, key, value, ttl)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock.ExpectSet(key, value, ttl).SetErr(errors.New("error"))
		err := r.Set(ctx, key, value, ttl)
		assert.Error(t, err)
	})
}

func TestGetRepository(t *testing.T) {
	client, mock := redismock.NewClientMock()
	r := &Repository{
		redis: client,
	}

	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock.ExpectGet(key).SetVal(value)
		result, err := r.Get(ctx, key)
		assert.NoError(t, err)
		assert.Equal(t, value, result)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock.ExpectGet(key).SetErr(errors.New("error"))
		_, err := r.Get(ctx, key)
		assert.Error(t, err)
	})
}

func TestDelRepository(t *testing.T) {
	client, mock := redismock.NewClientMock()
	r := &Repository{
		redis: client,
	}

	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock.ExpectDel(key).SetVal(int64(1))
		err := r.Del(ctx, key)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock.ExpectDel(key).SetErr(errors.New("error"))
		err := r.Del(ctx, key)
		assert.Error(t, err)
	})
}
