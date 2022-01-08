package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	var arrays [][]int
	for _, eachline := range fileTextLines {
		strNums := strings.Split(eachline, " ")
		var intNums []int
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			intNums = append(intNums, num)
		}
		arrays = append(arrays, intNums)
	}
	fmt.Printf("%d\n", findAdjLines(arrays, 4))

}
func findAdjLines(arrays [][]int, n int) int {
	subArrayLength := len(arrays[0])
	ArrayLength := len(arrays)
	var maxSum int = 0
	for arrayN := range arrays {
		for subArrayN := 0; subArrayN < subArrayLength; subArrayN++ {

			if subArrayN < subArrayLength-n+1 {
				horizontalAdjLineSum := findMultiplationOfArray(findHorizontallyAdjLine(arrays, subArrayN, arrayN, n))
				fmt.Printf("%d\n", horizontalAdjLineSum)

				if horizontalAdjLineSum > maxSum {
					maxSum = horizontalAdjLineSum

				}
			}
			if arrayN < ArrayLength-n+1 {
				verticalAdjLineSum := findMultiplationOfArray(findVerticallyAdjLine(arrays, subArrayN, arrayN, n))
				fmt.Printf("%d\n", verticalAdjLineSum)

				if verticalAdjLineSum > maxSum {
					maxSum = verticalAdjLineSum
				}
			}
			if arrayN < ArrayLength-n+1 && subArrayN < subArrayLength-n+1 {
				crossAdjLineSum := findMultiplationOfArray(findCrossLine(arrays, subArrayN, arrayN, n))
				fmt.Printf("%d\n", crossAdjLineSum)

				if crossAdjLineSum > maxSum {
					maxSum = crossAdjLineSum
				}
			}

		}
	}
	return maxSum
}
func findCrossLine(arrays [][]int, startX int, startY int, count int) []int {
	var crossLine []int
	for i := 0; i < count; i++ {
		num := arrays[startY+i][startX+i]
		crossLine = append(crossLine, num)
	}
	return crossLine

}
func findVerticallyAdjLine(arrays [][]int, startX int, startY int, count int) []int {
	var crossLine []int
	for i := 0; i < count; i++ {
		num := arrays[startY+i][startX]
		crossLine = append(crossLine, num)
	}
	return crossLine

}
func findHorizontallyAdjLine(arrays [][]int, startX int, startY int, count int) []int {
	var crossLine []int
	// if len(arrays)-1 > startY+count {
	// 	count = len(arrays) - 1 - startY
	// }
	for i := 0; i < count; i++ {
		num := arrays[startY][startX+i]
		crossLine = append(crossLine, num)
	}
	return crossLine

}
func findMultiplationOfArray(array []int) int {
	var sum int = 1

	for _, elem := range array {
		sum *= elem
	}
	return sum
}
