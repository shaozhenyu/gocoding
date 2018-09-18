package main

import (
	"fmt"
)

func main() {
	// n := []int{1,2,3,4,4,5,6,6,6,7,8,8,9,9,9,9,9,9}
	// n := []int{6,12,7}
	n := []int{39,19,30,56,79,50,19,70,30}
	fmt.Println(subarrayBitwiseORs(n))
}

func subarrayBitwiseORs(A []int) int {
	if len(A) == 0 {
		return 0
	}
	m := make(map[int]struct{})
	for i := 0; i < len(A); i++ {
		if i > 0 && A[i-1] == A[i] {
			continue
		}
		start := A[i]
		m[start] = struct{}{}
		for j := i+1; j < len(A); j++ {
			start = start | A[j]
			m[start] = struct{}{}
		}
	}
	return len(m)
}
