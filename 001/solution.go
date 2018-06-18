package e001

// https://projecteuler.net/problem=1
func solution(limit int) int {
	total := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}
