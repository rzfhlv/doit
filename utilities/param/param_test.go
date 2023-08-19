package param

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateOffset(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		expect := 0
		param := Param{
			Page:  1,
			Limit: 10,
			Total: 100,
		}
		offset := param.CalculateOffset()
		assert.Equal(t, expect, offset)
	})
}
