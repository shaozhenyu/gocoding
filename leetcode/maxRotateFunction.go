package main

import "fmt"

func main() {
	A := []int{4, 3, 2, 6}
	fmt.Println(maxRotateFunction(A))
}

func maxRotateFunction(A []int) int {
	totalSum := 0
	oneSum := 0
	for i := 0; i < len(A); i++ {
		totalSum += A[i] * i
		oneSum += A[i]
	}
	size := len(A)
	max := totalSum
	for i := len(A) - 1; i > 0; i-- {
		totalSum += oneSum - size * A[i]
		if totalSum > max {
			max = totalSum
		}
	}
	return max
}
