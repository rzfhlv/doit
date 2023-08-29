package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasher(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		password := "password"

		h := HasherPassword{}
		hashed, err := h.HashedPassword(password)

		assert.NoError(t, err)
		assert.NotNil(t, hashed)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		password := "passwordghhghgkgkggkjgjkgjkgkjgkgjkgggkjgkgjkgkjjkjljlkjkljljkljlkjljljklj"

		h := HasherPassword{}

		_, err := h.HashedPassword(password)
		assert.Error(t, err)
	})
}

func TestHasherValidate(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		hashed := "$2a$10$d3.zWWlz0tAnXis7fAJulumr2JHT5YDoZ7OzY9yJcx1TmQhS7c4mO"
		password := "password"

		h := HasherPassword{}
		err := h.VerifyPassword(hashed, password)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		hashed := "$2a$10$d3.zWWlz0tAnXis7fAJulumr2JHT5YDoZ7OzY9yJcx1TmQhS7c4mO"
		password := "invalidPassword"

		h := HasherPassword{}
		err := h.VerifyPassword(hashed, password)
		assert.Error(t, err)
	})
}
