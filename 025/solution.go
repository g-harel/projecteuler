package e025

import (
	"math/big"
)

// https://projecteuler.net/problem=25
func solution(size int) int {
	if size == 1 {
		return 1
	}

	a, b := new(big.Int).SetInt64(1), new(big.Int).SetInt64(1)
	for count := 3; true; count++ {
		a, b = b, new(big.Int).Add(a, b)
		if len(b.String()) >= size {
			return count
		}
	}

	return -1
}
