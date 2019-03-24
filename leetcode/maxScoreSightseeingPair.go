package main

import "fmt"

func main() {
	A := []int{8,1,9,5,2,6,20}
	fmt.Println(maxScoreSightseeingPair(A))
}

func maxScoreSightseeingPair(A []int) int {
	max := -2 << 31
	maxS := make([]int, len(A))
	for i := len(A) - 1; i > 0; i-- {
		v := A[0] + A[i] - i
		if v > max {
			max = v
		}
		maxS[i-1] = max
	}
	fmt.Println(maxS)
	for i := 1; i < len(A) - 1; i++ {
		maxS[i] += i + A[i] - A[0]
		if maxS[i] > max {
			max = maxS[i]
		}
	}
	fmt.Println(maxS)
	return max
}
