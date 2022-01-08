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

	var grid [][]int
	for _, eachline := range fileTextLines {
		strNums := strings.Split(eachline, " ")
		var intNums []int
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			intNums = append(intNums, num)
		}
		grid = append(grid, intNums)
	}
	fmt.Printf("%d\n", findAdjLines(grid, 4))

}
func findAdjLines(grid [][]int, n int) int {
	subArrayLength := len(grid[0])
	ArrayLength := len(grid)
	var maxSum int = 0

	for arrayN := range grid {
		for subArrayN := 0; subArrayN < subArrayLength; subArrayN++ {

			if subArrayN < subArrayLength-n+1 {
				horizontalAdjLineArr := findHorizontallyAdjLine(grid, subArrayN, arrayN, n)
				horizontalAdjLineSum := findMultiplationOfArray(horizontalAdjLineArr)
				fmt.Printf("horizontalAdjLineArr:%v\n", horizontalAdjLineArr)

				if horizontalAdjLineSum > maxSum {
					maxSum = horizontalAdjLineSum

				}
			}
			if arrayN < ArrayLength-n+1 {
				verticalAdjLineArr := findVerticallyAdjLine(grid, subArrayN, arrayN, n)
				verticalAdjLineSum := findMultiplationOfArray(verticalAdjLineArr)
				fmt.Printf("verticalAdjLineArr:%v\n", verticalAdjLineArr)

				if verticalAdjLineSum > maxSum {
					maxSum = verticalAdjLineSum
				}
			}
			if arrayN < ArrayLength-n+1 && subArrayN < subArrayLength-n+1 {
				rightCrossAdjLineArr := findRightCrossLine(grid, subArrayN, arrayN, n)
				crossAdjLineSum := findMultiplationOfArray(rightCrossAdjLineArr)
				fmt.Printf("findRightCrossLine:%v\n", rightCrossAdjLineArr)

				if crossAdjLineSum > maxSum {
					maxSum = crossAdjLineSum
				}
			}
			if subArrayN > n && arrayN < ArrayLength-n+1 {
				leftCrossAdjLineArr := findLeftCrossLine(grid, subArrayN, arrayN, n)
				leftCrossAdjLineSum := findMultiplationOfArray(leftCrossAdjLineArr)
				fmt.Printf("findLeftCrossLine:%v\n", leftCrossAdjLineArr)

				if leftCrossAdjLineSum > maxSum {
					maxSum = leftCrossAdjLineSum
				}
			}
			fmt.Printf("\n")

		}
	}
	return maxSum
}
func findRightCrossLine(arrays [][]int, startX int, startY int, count int) []int {
	var crossLine []int
	for i := 0; i < count; i++ {
		num := arrays[startY+i][startX+i]
		crossLine = append(crossLine, num)
	}
	return crossLine

}
func findLeftCrossLine(arrays [][]int, startX int, startY int, count int) []int {
	var crossLine []int
	for i := 0; i < count; i++ {
		num := arrays[startY+i][startX-i]
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
	sum := 1
	for _, elem := range array {
		sum *= elem
	}
	fmt.Printf("sum %d\n", sum)
	return int(sum)
}
