package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers []string
	for fileScanner.Scan() {
		numbers = append(numbers, fileScanner.Text())
	}

	var reverseDigitArray []int
	remainder := 0
	for currDigitPos := len(numbers[0]) - 1; currDigitPos >= 0; currDigitPos-- {
		sumOfDigits := 0
		for n := 0; n < len(numbers); n++ {
			digitChar := numbers[n][currDigitPos]
			digit := digitChar - 48

			sumOfDigits += int(digit)
		}
		remainder += sumOfDigits
		reverseDigitArray = append(reverseDigitArray, remainder%10)
		remainder /= 10
	}

	for remainder > 1 {
		reverseDigitArray = append(reverseDigitArray, remainder%10)
		remainder /= 10
	}
	digitArray := reverseArray(reverseDigitArray)
	fmt.Printf("%v\n", digitArray[:10])
}

func reverseArray(array []int) []int {

	var newArray []int
	for i := len(array) - 1; i > 1; i-- {
		newArray = append(newArray, array[i])
	}
	return newArray
}
