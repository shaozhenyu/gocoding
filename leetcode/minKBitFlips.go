package main

import "fmt"

func main() {
	fmt.Println(minKBitFlips([]int{0, 1, 1}, 1))
}

func minKBitFlips(A []int, K int) int {
	count := 0
	flip := 0
	change := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		flip ^= change[i]
		if A[i] == flip {
			count++
			if i + K > len(A) {
				return -1
			}
			flip ^= 1
			if i + K < len(A) {
				change[i + K] ^= 1
			}
		}
	}
	return count
}
