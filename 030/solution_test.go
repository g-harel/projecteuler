package e030

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	type Case struct {
		input  int
		result int
	}

	var cases = []Case{
		{0, 0},
		{1, 44},
		{2, 0},
		{3, 1301},
		{4, 19316},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, int64(c.result), solution(int64(c.input)), strconv.Itoa(c.input))
		}
	})
}
