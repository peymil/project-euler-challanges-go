package main

import "fmt"

func main() {
	fmt.Printf("%d", findSmallestNumDivisibleUntilM(20))
}
func isDivisibleUntilM(n uint32, m uint32) bool {
	var i uint32 = 2
	for i < m {
		if n%i != 0 {
			return false
		}
		i++
	}
	return true
}
func findSmallestNumDivisibleUntilM(m uint32) uint32 {
	var i uint32 = 20
	for {
		isEvenlyDivisible := isDivisibleUntilM(i, 20)
		if isEvenlyDivisible {
			return i
		}
		i++
	}
}
