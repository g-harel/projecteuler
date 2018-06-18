package projecteuler

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrime(t *testing.T) {
	type Case struct {
		input  int
		result bool
	}

	var cases = []Case{
		{9, false},
		{199, true},
		{207, false},
		{379, true},
		{1234567, false},
		{4562123, false},
		{1301081, true},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			assert.Equal(t, c.result, Prime(c.input), strconv.Itoa(c.input))
		}
	})
}
