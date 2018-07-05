package e020

import (
	"math/big"
	"strconv"
)

// https://projecteuler.net/problem=20
func solution(size int) int64 {
	fac := new(big.Int).MulRange(1, int64(size))
	bytes, err := fac.MarshalText()
	if err != nil {
		panic(err)
	}
	total := int64(0)
	for _, b := range bytes {
		digit, err := strconv.Atoi(string(b))
		if err != nil {
			panic(err)
		}
		total += int64(digit)
	}
	return total
}
