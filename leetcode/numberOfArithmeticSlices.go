package main

import (
	"fmt"
)

func main() {
	A := []int{79,20,64,28,67,81,60,58,97,85,92,96,82,89,46,50,15,2,36,44,54,2,90,37,7,79,26,40,34,67,64,28,60,89,46,31,9,95,43,19,47,64,48,95,80,31,47,19,72,99,28,46,13,9,64,4,68,74,50,28,69,94,93,3,80,78,23,80,43,49,77,18,68,28,13,61,34,44,80,70,55,85,0,37,93,40,47,47,45,23,26,74,45,67,34,20,33,71,48,96}
	A = []int{1,1,1,1,1,1}
	fmt.Println(numberOfArithmeticSlices(A))
}

func numberOfArithmeticSlices(A []int) int {
	if len(A) == 0 {
		return 0
	}
	m := make([]map[int]map[int]int, len(A))
	m[0] = make(map[int]map[int]int)
	ret := 0
	preAdd := 0
	for i := 1; i < len(A); i++ {
		m[i] = make(map[int]map[int]int) // key1: de key2: len
		j := 0
		if A[i] == A[i-1] {
			for k, v := range m[i-1] {
				m[i][k] = make(map[int]int)
				for k1, v1 := range v {
					m[i][k][k1] = v1
				}
			}
			ret += preAdd
			j = i - 1
		} else {
			preAdd = 0
		}
		fmt.Println(preAdd)
		for ; j < i; j++ {
			d := A[i] - A[j]
			if _, ok := m[i][d]; !ok {
				m[i][d] = make(map[int]int)
			}
			if p, ok := m[j][d]; ok {
				for k, v := range p {
					preAdd += v
					ret += v
					m[i][d][k+1] += v
				}
			}
			m[i][d][1] += 1
		}
		fmt.Println(i, ret)
		fmt.Println(m[i])
	}
	// fmt.Println(m)
	return ret
}
