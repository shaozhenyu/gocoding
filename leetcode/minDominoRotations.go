package main

import "fmt"

func main() {
	A := []int{2,1,2,4,2,2}
	B := []int{5,2,6,2,3,2}
	fmt.Println(minDominoRotations(A, B), 2)
}

func minDominoRotations(A []int, B []int) int {
	nums := [7]int{}
	as := [7]int{}
	bs := [7]int{}
	for i := 0; i < len(A); i++ {
		if A[i] == B[i] {
			nums[A[i]]++
		} else {
			nums[A[i]]++
			as[A[i]]++
			nums[B[i]]++
			bs[B[i]]++
		}
	}
	fmt.Println(nums)
	fmt.Println(as)
	fmt.Println(bs)

	for k, v := range nums {
		if v == len(A) {
			return min(as[k], bs[k])
		}
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
