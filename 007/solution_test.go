package e001

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
		{1, 2},
		{2, 3},
		{6, 13},
		{191, 1153},
		{757, 5749},
		{10001, 104743},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, c.result, solution(c.input), strconv.Itoa(c.input))
		}
	})
}
