package main

import "fmt"

func main() {
	fmt.Println(minScoreTriangulation([]int{1,2,3}))
}

func minScoreTriangulation(A []int) int {
	w := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		w[i] = make([]int, len(A))
	}
	for length := 2; length < len(A); length++ {
		for i := 0; i < len(A) - length; i++ {
			j := i + length
			min := 2 << 31
			for k := i + 1; k < j; k++ {
				now := w[i][k] + w[k][j] + A[i] * A[k] * A[j]
				if now < min {
					min = now
				}
			}
			w[i][j] = min
		}
	}
	return w[0][len(A)-1]
}
