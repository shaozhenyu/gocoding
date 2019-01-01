package main

import "fmt"

func main() {
	A := []int{9,8,1,0,1,9,4,0,4,1}
	fmt.Println(maxWidthRamp(A))
}

func maxWidthRamp(A []int) int {
	max := 0
	for i := len(A) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if i - j  < max {
				break
			}
			if A[j] <= A[i] {
				max = i - j
				break
			}
		}
	}
	return max
}
