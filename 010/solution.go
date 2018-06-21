package e010

import "github.com/g-harel/projecteuler"

// https://projecteuler.net/problem=10
func solution(limit int) int {
	total := 0
	for i, prime := range projecteuler.Primes(limit) {
		if prime {
			total += i
		}
	}
	return total
}
