package main

import "fmt"

func main() {

}

func repeatedNTimes(A []int) int {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		if m[A[i]] > 1 {
			return A[i]
		}
		m[A[i]]++
	}
	return -1
}
