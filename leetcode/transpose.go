package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{[]int{1,2,3},[]int{4,5,6}}))
}

func transpose(A [][]int) [][]int {
	B := make([][]int, len(A[0]))
	for i := 0; i < len(B); i++ {
		B[i] = make([]int, len(A))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			B[j][i] = A[i][j]
		}
	}
	return B
}
