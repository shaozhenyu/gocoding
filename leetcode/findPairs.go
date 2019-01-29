package main

import "fmt"

func main() {
	nums := []int{3, 1, 4, 1, 5}
	k := 2
	fmt.Println(findPairs(nums, k))
}

func findPairs(nums []int, k int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	ret := 0
	for num, count := range m {
		if k == 0 {
			if count > 1 {
				ret++
			}
		} else {
			if _, ok := m[num + k]; ok {
				ret++
			}
		}
	}
	return ret
}
