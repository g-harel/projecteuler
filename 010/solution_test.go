package e010

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
		{10, 17},
		{754, 45434},
		{91291, 381349157},
		{2000000, 142913828922},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, c.result, solution(c.input), strconv.Itoa(c.input))
		}
	})
}
