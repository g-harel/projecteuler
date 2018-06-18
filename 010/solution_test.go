package e001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("should compute the correct value for 10", func(t *testing.T) {
		assert.Equal(t, 17, solution(10), "")
	})

	t.Run("should compute the correct value for 2000000", func(t *testing.T) {
		assert.Equal(t, 142913828922, solution(2000000), "")
	})
}
