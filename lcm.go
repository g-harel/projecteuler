package projecteuler

// LCM computes the lowest common multiple.
func LCM(list ...int) int {
	if len(list) < 2 {
		panic("list is too short")
	}

	a := list[0]
	for _, num := range list[1:] {
		a = lcm(a, num)
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
