package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%d\n", findHighlyDivableTriNum(500))
}

func findHighlyDivableTriNum(maxDivisorCount int) int {
	currMaxDivisorCount := 0
	i := 1
	currSum := 0
	for currMaxDivisorCount <= maxDivisorCount {
		currSum += i
		currDivisiorCount := findDivisorCount(currSum)
		if currMaxDivisorCount < currDivisiorCount {
			currMaxDivisorCount = currDivisiorCount
			fmt.Printf("%d %d\n", currMaxDivisorCount, currSum)
		}
		i++
	}
	return currSum
}

func findDivisorCount(number int) int {
	divisorCount := 0
	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			divisorCount++
		}
	}
	return divisorCount * 2

}
