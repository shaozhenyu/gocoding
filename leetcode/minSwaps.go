package main

import "fmt"

func main() {
	A := []int{1, 2, 5, 4, 7}
	B := []int{1 ,2, 3, 6, 6}
	fmt.Println(minSwap(A, B))
}

func minSwap(A []int, B []int) int {
	keep := make([]int, len(A))
	swap := make([]int, len(A))
	swap[0] = 1
	for i := 1; i < len(A); i++ {
		keep[i], swap[i] = len(A), len(A)
		if A[i-1] < A[i] && B[i-1] < B[i] {
			keep[i] = keep[i-1]
			swap[i] = swap[i-1] + 1
		}
		if A[i-1] < B[i] && B[i-1] < A[i] {
			keep[i] = min(keep[i], swap[i-1])
			swap[i] = min(swap[i], keep[i-1] + 1)
		}
	}
	fmt.Println(keep)
	fmt.Println(swap)
	return min(keep[len(A)-1], swap[len(A)-1])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
