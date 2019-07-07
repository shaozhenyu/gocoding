package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0}, 2))
}

func longestOnes(A []int, K int) int {
	max := -1
	start := 0
	stack := make([]int, len(A))
	startIdx, endIdx := 0, 0
	for i := 0; i< len(A); i++ {
		if A[i] == 0 {
			stack[endIdx] = i
			endIdx++
			if K > 0 {
				K--
			} else {
				start = stack[startIdx] + 1
				startIdx++
			}
		}
		if i - start + 1 > max {
			max = i - start + 1
		}
	}
	return max
}
