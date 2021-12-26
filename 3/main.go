package main

//https://en.wikipedia.org/wiki/Pollard's_rho_algorithm

import "fmt"

func main() {
	fmt.Printf("%d", findHighestFactor(600851475143))
}
func findHighestFactor(n uint64) uint64 {
	var highestPrimeFactor uint64 = 0
	var d uint64 = 3
	for n%2 == 0 {
		//the only odd prime number is 2.
		n /= 2
		if highestPrimeFactor < d {
			highestPrimeFactor = d
		}
	}
	for d < n {
		for n%d == 0 {
			n /= d
			if highestPrimeFactor < d {
				highestPrimeFactor = d
			}
		}
		//proceed by adding 2 to d. There is no need to check even numbers.
		d += 2
	}
	if n > highestPrimeFactor {
		//if final n is greater than current highestPrimeFactor
		return n
	}
	return highestPrimeFactor
}
