package main

import "fmt"

func main() {
	fmt.Println(smallestRepunitDivByK(12345679))
}


func smallestRepunitDivByK(K int) int {
	if K%2 == 0 {
		return -1
	}
	sum := 0
	for i := 1; i <= K * 10; i++ {
		sum = (sum * 10 + 1)%K
		if sum == 0 {
			return i
		}
	}
	return -1
}
