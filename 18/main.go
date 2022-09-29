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
	var tri [][]int
	for fileScanner.Scan() {
		scannedStrings := strings.Split(fileScanner.Text(), " ")
		var arrToAppend []int
		for _, strNum := range scannedStrings {

			intVar, _ := strconv.Atoi(strNum)
			arrToAppend = append(arrToAppend, intVar)

		}
		tri = append(tri, arrToAppend)
	}
	fmt.Printf("%d\n", findHighestChainSumInTri(tri))

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findHighestChainSumInTri(arr [][]int) int {
	for i := len(arr) - 2; i >= 0; i-- {
		for a := 0; a < len(arr[i]); a++ {
			highest := max(arr[i+1][a], arr[i+1][a+1])
			arr[i][a] = arr[i][a] + highest
		}
		fmt.Printf("[%v]\n", arr[i])

	}
	return arr[0][0]
}
