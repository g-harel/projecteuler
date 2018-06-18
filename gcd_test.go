package projecteuler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	type Case struct {
		inputs []int
		result int
	}

	var cases = []Case{
		{[]int{2, 4}, 2},
		{[]int{6, 9}, 3},
		{[]int{35, 20}, 5},
		{[]int{19, 19, 19}, 19},
		{[]int{14, 21, 49}, 7},
		{[]int{25, 90, 135, 140}, 5},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			b, _ := json.Marshal(c.inputs)
			assert.Equal(t, c.result, GCD(c.inputs...), string(b))
		}
	})
}
