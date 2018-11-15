package main

import "fmt"

func main() {
	A := []int{0}
	fmt.Println(isIdealPermutation(A))
}
// i < j A[i] > A[j]
//  A[i] > A[i+1]
func isIdealPermutation(A []int) bool {
	if len(A) == 0 {
		return false
	}
	global := A[0]
	local := 0
	max, min := A[0], A[0]
	for i := 1; i < len(A); i++ {
		if A[i-1] > A[i] {
			local++
		}
		if A[i] > max {
			global += A[i] - i
			max = A[i]
			continue
		}
		if A[i] < min {
			global += A[i]
			min = A[i]
			continue
		}
		for j := 0; A[j] != i; j++ {
			if A[j] > i {
				global++
			}
		}
	}
	return global == local
}
