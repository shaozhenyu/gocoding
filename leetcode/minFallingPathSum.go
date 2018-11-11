package main

import "fmt"

func main() {
	A := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(minFallingPathSum(A))
}

func minFallingPathSum(A [][]int) int {
	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] += getMinPre(i, j, A)
		}
	}
	ret := 2 << 31
	last := len(A) - 1
	for j := 0; j < len(A[0]); j++ {
		ret = min(A[last][j], ret)
	}
	return ret
}

func getMinPre(i, j int, A [][]int) int {
	ret := A[i-1][j]
	if j > 0 {
		ret = min(ret, A[i-1][j-1])
	}
	if j < len(A[i])-1 {
		ret = min(ret, A[i-1][j+1])
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
