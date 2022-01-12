package main

import "fmt"

var countingNumbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var afterHundred = []string{"hundred", "thousand"}
var afterTens = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var tensList = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

func main() {
	fmt.Printf("%d", len(countNumLettersUntilN(1000)))
	// fmt.Printf("%s", numToString(1000))
}
func countNumLettersUntilN(n int) string {
	var fullString string
	for i := 1; i <= n; i++ {
		fullString += numToString(i)
	}
	return fullString
}

func numToString(n int) string {
	var resultStr string
	if n >= 100 {
		c := 0
		for i := n / 100; i >= 1; i /= 10 {
			a := i % 10

			if a != 0 {
				resultStr = countingNumbers[a] + afterHundred[c] + resultStr
				if n%100 != 0 {
					resultStr += "and"
				}
			}

			c++
		}
	}
	tens := (n % 100) / 10
	ones := n % 10

	if tens == 1 {
		resultStr += tensList[ones]
	} else {
		if tens >= 2 {
			resultStr += afterTens[tens-2]
		}
		if ones == 0 {
			resultStr += ""
		} else {
			resultStr += countingNumbers[ones]
		}
	}
	return resultStr
}
