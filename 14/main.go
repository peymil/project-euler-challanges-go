package main

import "fmt"

// n → n/2 (n is even)
// n → 3n + 1 (n is odd)

func main() {
	fmt.Printf("%v", getFirstNumOfLongestCollatzChain(1000000))
}
func getNextCollatzSeqMember(num uint) uint {
	n := uint(0)
	//if is even
	if num%2 == 0 {
		n = num / 2
	} else {
		n = num*3 + 1
	}
	return n
}
func getCollatzSeqCount(startCount uint) uint {
	count := uint(0)
	for currNum := startCount; currNum > 1; currNum = getNextCollatzSeqMember(currNum) {
		count++
	}
	return count
}
func getFirstNumOfLongestCollatzChain(max uint) uint {
	longestsChainCount := uint(0)
	longestsChainStartNumber := uint(0)
	for i := uint(1); i < max; i++ {
		chainCount := getCollatzSeqCount(i)
		if chainCount > longestsChainCount {
			longestsChainCount = chainCount
			longestsChainStartNumber = i
		}
	}
	return longestsChainStartNumber
}
