package main

import "fmt"

func main() {
	squareOfSum, sumOfSquares := sumSquareDiffOfNNumber(100)
	fmt.Printf("%d", squareOfSum-sumOfSquares)
}
func sumSquareDiffOfNNumber(n int32) (int32, int32) {
	var i int32 = 1
	var sumOfSquares int32 = 0
	var sum int32 = 0

	for i <= n {
		fmt.Printf("%d\n", i)
		sumOfSquares += i * i
		sum += i
		i++
	}
	return sum * sum, sumOfSquares
}
