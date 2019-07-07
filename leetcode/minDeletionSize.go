package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"babca","bbazb"}), 3)
	fmt.Println(minDeletionSize([]string{"cbbdabc"}), 3)
}

func minDeletionSize(A []string) int {
	sort := make([]int, len(A[0]))
	for i := 0; i < len(sort); i++ {
		sort[i] = 1
	}
	max := 1
	for i := 1; i < len(A[0]); i++ {
		for j := 0; j < i; j++ {
			for k := 0; k <= len(A); k++ {
				if k == len(A) {
					sort[i] = Max(sort[i], sort[j] + 1)
					max = Max(max, sort[i])
				} else if A[k][j] > A[k][i] {
					break
				}
			}
		}
	}
	fmt.Println(sort)
	return len(A[0]) - max
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
