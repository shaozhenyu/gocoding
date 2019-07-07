package main

import (
	"fmt"
)

func main() {
	A := []int{10, 9, 3, 2}
	k := 2
	fmt.Println(maxSumAfterPartitioning(A, k))
}

func maxSumAfterPartitioning(A []int, K int) int {
	sum := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		end := i + K - 1
		if end >= len(A) {
			end = len(A) - 1
		}
		// fmt.Println(i, end)
		for j := i; j <= end; j++ {
			start := j - K + 1
			if start < 0 {
				start = 0
			}
			// fmt.Println("xx: ", start, j)
			for t := start; t <= i; t++ {
				pre := 0
				if t - 1 >= 0 {
					pre = sum[t-1]
				}
				// fmt.Println("sum: ", pre, A[i] * (j - t + 1))
				sum[j] = Max(sum[j], pre + A[i] * (j - t + 1))
				fmt.Println(sum)
			}
		}
	}
	return sum[len(A)-1]
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}


