package projecteuler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCM(t *testing.T) {
	type Case struct {
		inputs []int
		result int
	}

	var cases = []Case{
		{[]int{2, 4}, 4},
		{[]int{12, 32}, 96},
		{[]int{2, 3, 9}, 18},
		{[]int{7, 7, 7}, 7},
		{[]int{10, 25, 45, 60}, 900},
	}

	t.Run("should compute the correct value for all cases", func(t *testing.T) {
		for _, c := range cases {
			b, _ := json.Marshal(c.inputs)
			assert.Equal(t, c.result, LCM(c.inputs...), string(b))
		}
	})
}
