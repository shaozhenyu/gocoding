package main

import (
	"fmt"
)

func main() {
	a := []int{0,0,0}
	fmt.Println(countTriplets(a))
}

func countTriplets(A []int) int {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		m[i]++
	}
	mc := make(map[int]int)
	count := 0
	for i := 0; i < len(A); i++ {
		if _, ok := mc[A[i]]; ok {
			count += mc[A[i]]
			continue
		}
		c := 0
		for j := 0; j < len(A); j++ {
			x := A[i] & A[j]
			if x == 0 {
				c += len(A)
			} else {
				for k := 0; k < len(A); k++ {
					if x & A[k] == 0 {
						c++
					}
				}
			}
		}
		count += c
		mc[A[i]] = c
	}
	return count
}
