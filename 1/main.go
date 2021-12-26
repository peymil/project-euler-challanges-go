package main

import "fmt"

func main() {
	sumOfDivables := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sumOfDivables += i
		}
	}
	fmt.Printf("%d", sumOfDivables)
}
