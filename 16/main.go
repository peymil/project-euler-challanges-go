package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Printf("%d", findSumOfPowDigits(2, 1000))

}
func findSumOfPowDigits(n int64, pow int64) uint64 {
	x, y := big.NewInt(n), big.NewInt(pow)
	str := x.Exp(x, y, nil).String()
	sum := uint64(0)
	for _, str := range str {
		num, _ := strconv.ParseUint(string(str), 10, 32)
		sum += num
	}
	return sum
}
