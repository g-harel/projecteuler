package e025

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
		{3, 12},
		{10, 45},
		{32, 151},
		{100, 476},
		{670, 3203},
		{1000, 4782},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, c.result, solution(c.input), strconv.Itoa(c.input))
		}
	})
}
