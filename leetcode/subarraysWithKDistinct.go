package main

import "fmt"

func main() {
	A := []int{1,2,1,3,4}
	K := 3
	fmt.Println(subarraysWithKDistinct(A, K))
}

func subarraysWithKDistinct(A []int, K int) int {
	m := make(map[int]int)
	start, end := 0, 0
	count := 0
	for start < len(A) {
		if end < len(A) {
			m[A[end]]++
		}
		if len(m) < K && end < len(A) {
			end++
		} else if len(m) == K && end < len(A) {
			// count++
			end++
		} else {
			if end < len(A) {
				m[A[end]]--
				if m[A[end]] <= 0 {
					delete(m, A[end])
				}
			}

			m[A[start]]--
			if m[A[start]] <= 0 {
				delete(m, A[start])
			}
			start++
			if len(m) == K {
				count += end - start + 1
			}
			if end < start {
				end = start
			}
		}
	}
	return count
}
