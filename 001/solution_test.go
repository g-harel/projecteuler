package e001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("should compute the correct value for 10", func(t *testing.T) {
		assert.Equal(t, 23, solution(10), "")
	})

	t.Run("should compute the correct value for 1000", func(t *testing.T) {
		assert.Equal(t, 233168, solution(1000), "")
	})
}
