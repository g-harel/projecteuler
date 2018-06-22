package e015

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
		{2, 6},
		{20, 1},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, c.result, solution(c.input), strconv.Itoa(c.input))
		}
	})
}
