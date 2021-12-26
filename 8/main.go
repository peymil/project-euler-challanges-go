package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	strContent := strings.ReplaceAll(string(content), "\n", "")

	fmt.Printf("%d\n", findHighestNAdjacentDigit(strContent, 13))
}
func findHighestNAdjacentDigit(number string, digitN uint32) uint64 {

	iterateNTimes := uint32(len(number)) - digitN
	var highestNumber uint64 = 0
	for i := uint32(0); i <= iterateNTimes; i++ {
		currNumSlice := number[i : digitN+i]
		var product uint64 = 1
		for _, num := range currNumSlice {
			// digit char to int
			product = uint64(num-48) * product
		}
		if product > highestNumber {
			highestNumber = product
		}
	}
	return highestNumber
}
