package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2,2,1}))
}

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 3; i >= 0; i-- {
		if A[i+2] < A[i+1] + A[i] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
