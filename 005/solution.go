package e001

import "github.com/g-harel/projecteuler"

// https://projecteuler.net/problem=5
func solution(limit int) int {
	list := make([]int, limit)
	for i := range list {
		list[i] = i + 1
	}
	return projecteuler.LCM(list...)
}
