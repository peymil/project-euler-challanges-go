package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Printf("%d", getNthPasTriIndex(20))
}

// https://stemhash.com/counting-lattice-paths/
// counting using pascal triangle
// In the triangle above, the line crosses the terms 1, 2, 6, 20, and 70. If we assign an index to these terms–starting from zero–the 4th term is 70.
// That is the number of lattice paths in a 4x4 lattice. In a 2x2 lattice, the number of paths is equal to the 2nd crossed-out term, which is 6.
// Just as we found at the beginning, with our 2x2 lattice example.

func getNthPasTriIndex(n int64) *big.Int {
	realN := n * 2
	z := new(big.Int)
	return z.Binomial(realN, realN/2)
}

// func factorial(n uint64) uint64 {
// 	result := uint64(1)
// 	for i := uint64(1); i <= n; i++ {
// 		result *= i
// 	}
// 	return result
// }

// n!/(k!(n-k)!)
// func binomialCoefficient(n uint64, k uint64) uint64 {
// 	return factorial(n) / factorial(k) * factorial(n-k)
// }
