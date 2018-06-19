package e001

import (
	"github.com/g-harel/projecteuler"
)

// https://projecteuler.net/problem=10
func solution(nth int) int {
	limit := nth
	for {
		primes := projecteuler.Primes(limit * 2)
		count := 0
		for i, prime := range primes {
			if prime {
				count++
			}
			if count == nth {
				return i
			}
		}
		limit = len(primes)
	}
}
