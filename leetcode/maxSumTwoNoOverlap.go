package main

import "fmt"

func main() {
	A := []int{0,6,5,2,2,5,1,9,4}
	A = []int{1,2,3}
	L := 2
	M := 1
	fmt.Println(maxSumTwoNoOverlap(A, L, M))
}

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	return Max(mst(A, L, M), mst(A, M, L))
}

func mst(A []int, L int, M int) int {
	maxL := make([]int, len(A)+1)
	sumL := 0
	i := len(A) - 1
	for ; i > len(A) - L; i-- {
		sumL += A[i]
	}
	end := len(A) - 1
	for i >= 0 {
		sumL = sumL + A[i]
		maxL[i] = Max(sumL, maxL[i+1])
		i--
		sumL -= A[end]
		end--
	}
	i = 0
	sumM := 0
	max := 0
	for ; i < M - 1; i++ {
		sumM += A[i]
	}
	start := 0
	for ; i < len(A) - L; i++ {
		sumM += A[i]
		max = Max(sumM + maxL[i+1], max)
		sumM -= A[start]
		start++
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
