package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaker(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		f := FakerGenerator{}
		name := f.GenerateName()
		assert.NotNil(t, name)
	})
}
