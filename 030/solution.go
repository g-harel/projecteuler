package e030

import (
	"math/big"
)

func upper(pow int64) *big.Int {
	nine := new(big.Int).SetInt64(9)
	ten := new(big.Int).SetInt64(10)

	maxDigit := new(big.Int).Exp(
		nine,
		new(big.Int).SetInt64(pow),
		nil,
	)

	curr := nine
	for i := int64(1); true; i++ {
		I := new(big.Int).SetInt64(i)
		maxSum := new(big.Int).Mul(maxDigit, I)
		if maxSum.Cmp(curr) == -1 {
			return curr
		}
		curr = new(big.Int).Mul(curr, ten)
		curr = new(big.Int).Add(curr, nine)
	}

	return new(big.Int).SetInt64(-1)
}

func sum(digits *big.Int, pow int64) *big.Int {
	zero := new(big.Int).SetInt64(0)
	ten := new(big.Int).SetInt64(10)
	exp := new(big.Int).SetInt64(pow)

	sum := zero
	for digits.Cmp(zero) == 1 {
		digit := new(big.Int)
		digits, _ = new(big.Int).QuoRem(digits, ten, digit)
		val := new(big.Int).Exp(digit, exp, nil)
		sum = new(big.Int).Add(sum, val)
	}

	return sum
}

// https://projecteuler.net/problem=30
func solution(size int64) int64 {
	if size == 0 {
		return 0
	}

	one := new(big.Int).SetInt64(1)
	max := upper(size)

	total := new(big.Int).SetInt64(-1)
	i := new(big.Int).SetInt64(0)
	for i.Cmp(max) == -1 {
		if i.Cmp(sum(i, size)) == 0 {
			total = new(big.Int).Add(total, i)
		}
		i = new(big.Int).Add(i, one)
	}

	if !total.IsInt64() {
		return -1
	}
	return total.Int64()
}
