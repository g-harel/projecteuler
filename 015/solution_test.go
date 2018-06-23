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
		{7, 3432},
		{16, 601080390},
		{20, 137846528820},
		{33, 7219428434016265740},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, int64(c.result), solution(c.input), strconv.Itoa(c.input))
		}
	})
}
