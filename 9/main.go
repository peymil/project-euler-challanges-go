package main

import "fmt"

func main() {
	a, b, c := idontknow()
	fmt.Printf("%d,%d,%d\n", a, b, c)
	fmt.Printf("%d\n", a*b*c)
}
func idontknow() (int, int, int) {

	total := 0
	for {
		total++

		c := total % 1000
		b := (total / 1000) % 1000
		a := total / 1000000

		if a+b+c == 1000 && c*c+b*b == a*a && a+b+c != 0 {
			return a, b, c
		}
		if a == 1000 {
			fmt.Printf("bruhhh")
			return 0, 0, 0
		}
	}
}

// func sumOfArray(arr []int) int {
// 	sum := 0
// 	for _, number := range arr {
// 		sum += number
// 	}
// 	return sum
// }
