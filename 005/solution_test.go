package e001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("should compute the correct value for 10", func(t *testing.T) {
		assert.Equal(t, 2520, solution(10), "")
	})

	t.Run("should compute the correct value for 20", func(t *testing.T) {
		assert.Equal(t, 232792560, solution(20), "")
	})
}
