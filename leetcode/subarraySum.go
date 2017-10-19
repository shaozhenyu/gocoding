package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 1, 2, 3, 1, 2}
	k := 5
	fmt.Println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	sum := 0
	count := 0
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			count += m[sum-k]
		}

		if _, ok := m[sum]; ok {
			m[sum] += 1
		} else {
			m[sum] = 1
		}
	}
	return count
}
