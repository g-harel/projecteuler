package projecteuler

import (
	"math"
)

// Prime computes if the input is prime.
func Prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Primes computes the primality of all numbers until n.
func Primes(n int) []bool {
	if n < 0 {
		panic("number is negative")
	}
	if n == 0 {
		return []bool{false}
	}

	// create a slice to contain all values
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false

	// mark multiples in sieve
	for i, prime := range sieve {
		if !prime {
			continue
		}
		for j := i * i; j < len(sieve); j += i {
			sieve[j] = false
		}
	}

	return sieve
}

// PrimeN computes the nth prime.
func PrimeN(n int) int {
	if n < 1 {
		panic("number is too small")
	}
	if n < 6 {
		return []int{-1, 2, 3, 5, 7, 11}[n]
	}

	// compute the upwards bound of the nth prime
	nF := float64(n)
	max := nF*math.Log(nF) + nF*math.Log(math.Log(nF))

	// find nth prime
	primes := Primes(int(max))
	count := 0
	for i, prime := range primes {
		if prime {
			count++
			if count == n {
				return i
			}
		}
	}
	return -1
}
