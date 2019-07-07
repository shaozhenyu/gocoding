package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{1,2,3}, 3))
}

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	idx := 0
	for idx < len(A) && A[idx] < 0 && K > 0 {
		A[idx] = -1 * A[idx]
		K--
		idx++
	}
	if K > 0 && K%2 == 1 {
		sort.Ints(A)
		A[0] = -1 * A[0]
	}
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	return sum
}
