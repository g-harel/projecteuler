package projecteuler

// Prime computes if the input is prime.
func Prime(n int) bool {
	return Primes(n)[n]
}

// Primes computes the primality of all numbers until n and counts them.
func Primes(n int) []bool {
	if n < 0 {
		panic("number is negative")
	}

	seq := make([]bool, n+1)
	for i := range seq {
		seq[i] = true
	}
	seq[0], seq[1] = false, false

	for i, prime := range seq {
		if !prime {
			continue
		}
		for j := i + i; j < len(seq); j += i {
			seq[j] = false
		}
	}

	return seq
}
