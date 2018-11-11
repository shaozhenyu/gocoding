package main

import "fmt"

func main() {
	//A := []int{1, 2, 3, 5, 6, 8, 10, 11, 13, 17, 28, 45}
	// A := []int{2, 4, 5, 6, 7, 8, 11, 13, 14, 15, 21, 22, 34}
	A := []int{2,4,7,8,9,10,14,15,18,23,32,50}
	fmt.Println(lenLongestFibSubseq(A))
}

func lenLongestFibSubseq(A []int) int {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		m[A[i]] = 1
	}
	done := make(map[string]bool)
	ret := 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			a, b := A[i], A[j]
			if done[fmt.Sprintf("%d+%d", a, b)] {
				continue
			}
			count := 2
			for m[a+b] > 0 {
				done[fmt.Sprintf("%d+%d", a, b)] = true
				count++
				a, b = b, a+b
				if count > ret {
					ret = count
				}
			}
		}
	}
	return ret
}

func lenLongestFibSubseq1(A []int) int {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		m[A[i]] = 1
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			sum := A[i] + A[j]
			if _, ok := m[sum]; ok {
				if m[A[i]] == 1 || m[A[j]] == 1 {
					m[sum] = 3
				} else if m[A[i]]+1 == m[A[j]] {
					m[sum] = max(m[A[j]]+1, m[sum])
				} else {
					m[sum] = max(m[A[i]]+2, m[sum])
				}
			}
		}
	}
	fmt.Println(m)
	ret := 0
	for _, v := range m {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
