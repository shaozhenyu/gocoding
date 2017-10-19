package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(nums))
}

func findLHS(nums []int) int {
	m := make(map[int]int)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num] += 1
		} else {
			m[num] = 1
		}
	}

	lenN := len(nums)/2 + 1

	fmt.Println(m)
	max := -1
	for num, times := range m {

		if v, ok := m[num+1]; ok {
			sum := v + times
			if max < sum {
				max = sum
			}
		}

		if max > lenN {
			return max
		}
	}
	return max
}
