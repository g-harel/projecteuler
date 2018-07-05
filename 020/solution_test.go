package e020

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
		{10, 27},
		{37, 153},
		{63, 333},
		{100, 648},
		{176, 1215},
		{2490, 30501},
		{21309, 350352},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, int64(c.result), solution(c.input), strconv.Itoa(c.input))
		}
	})
}
