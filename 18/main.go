package main

import "fmt"

func main() {
	fmt.Printf("%v")
}

func findHighestChainSumInTri(arr [][]int) {
	for n := 0; n < len(arr)-2; n++ {
		higestNum, pos := findHighestNumInArr(arr[n+1])
		higestNum
		maxValueOfCurrentLine := 0
		arr[n]
	}
}
func findHighestNumInArr(nums []int) (int, int) {
	highestNum := 0
	highestNumPos := 0
	for pos, num := range nums {
		if num > highestNum {
			highestNum, highestNumPos = highestNum, pos
		}
	}
	return highestNum, highestNumPos
}
