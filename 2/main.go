package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d", findEvenFibs(1, 2, 4000000))
}
func findEvenFibs(a uint32, b uint32, max uint32) uint32 {
	var sum uint32
	//Check if inputs are even
	if a%2 == 0 {
		sum += a
	}
	if b%2 == 0 {
		sum += b
	}
	//
	for b < max {
		result := classicalFib(a, b)
		if result%2 == 0 {
			sum += result
		}
		a = b
		b = result
	}
	return sum
}

//Xn = [φn – (1-φ)n]/√5
func fib(n float64) float64 {
	return (math.Pow(math.Phi, n) - math.Pow((1-math.Phi), n)) / math.Sqrt(5)
}
func classicalFib(a uint32, b uint32) uint32 {
	return a + b
}
