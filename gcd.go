package projecteuler

// GCD computes the greates common divisor.
func GCD(list ...int) int {
	if len(list) < 2 {
		panic("list is too short")
	}

	a := list[0]
	for _, num := range list[1:] {
		a = gcd(a, num)
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
