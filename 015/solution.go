package e015

import (
	"math/big"
)

// https://projecteuler.net/problem=15
func solution(size int) int64 {
	nominator := new(big.Int).MulRange(int64(size)+1, int64(size)*2)
	denominator := new(big.Int).MulRange(1, int64(size))

	result := new(big.Float).Quo(
		new(big.Float).SetInt(nominator),
		new(big.Float).SetInt(denominator),
	)

	val, acuracy := result.Int64()
	if acuracy != 0 {
		return -1
	}
	return val
}
