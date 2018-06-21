package e005

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
		{2, 2},
		{10, 2520},
		{20, 232792560},
		{30, 2329089562800},
		{42, 219060189739591200},
		{98, 753127095974807136},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, c.result, solution(c.input), strconv.Itoa(c.input))
		}
	})
}
