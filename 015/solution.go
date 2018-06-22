package e015

import (
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=15
func solution(size int) int {
	bigTwo := new(big.Int).SetInt64(2)
	facSingle := new(big.Int).MulRange(1, int64(size))   // (size)!
	facDouble := new(big.Int).MulRange(1, int64(size)*2) // (2*size)!
	doubleFac := new(big.Int).Mul(facSingle, bigTwo)     // 2(size)!

	facDoubleF := new(big.Float).SetInt(facDouble)
	doubleFacF := new(big.Float).SetInt(doubleFac)
	result := new(big.Float).Quo(facDoubleF, doubleFacF) // (2*size)!/2(size)!
	b, _ := result.MarshalText()
	val, a := result.Int64()
	fmt.Println(a, string(b), val, int(val))
	return int(val)
}

/*

(2(a))!/a!a!

0011
0101
0110
1001
1010
1100

*/
