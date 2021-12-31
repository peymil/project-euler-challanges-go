package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d\n", findNthPrimeNumber(2000000))
}
func findNthPrimeNumber(n uint64) uint64 {
	var currNum uint64 = 3
	var sum uint64 = 2
	for {

		if isFib(currNum) {
			sum += currNum
		}
		currNum += 2

		if currNum >= n {
			break
		}
	}
	return sum
}
func isFib(n uint64) bool {
	if n == 2 {
		return true
	}
	for i := uint64(2); i <= uint64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
