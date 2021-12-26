package main

import "fmt"

func main() {
	fmt.Printf("%d", findNthPrimeNumber(10001))
}
func findNthPrimeNumber(n uint32) uint32 {
	var currNum uint32 = 3
	var primeN uint32 = 1
	for n > primeN {
		if isFib(currNum) {
			primeN++
		}
		currNum += 2
	}
	// discard +=2 on loop exit
	return currNum - 2
}
func isFib(n uint32) bool {
	if n == 2 {
		return true
	}
	for i := uint32(2); i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
