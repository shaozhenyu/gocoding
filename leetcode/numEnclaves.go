package main

import "fmt"

func main() {

}

func numEnclaves(A [][]int) int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				if i == 0 || j == 0 || i == len(A) - 1 || j == len(A[i]) - 1 {
					fly(i, j)
				}
			}
		}
	}
	count := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func fly(i, j int, A [][]int) {
	if A[i][j] == 0 {
		return
	}
	A[i][j] = 0
	if i > 0 {
		fly(i-1, j, A)
	}
	if i < len(A) - 1 {
		fly(i+1, j, A)
	}
	if j > 0 {
		fly(i, j - 1, A)
	}
	if j < len(A[i]) - 1 {
		fly(i, j + 1, A)
	}
}
