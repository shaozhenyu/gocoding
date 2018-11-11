package main

import "fmt"

func main() {
	nums := []int{1,1,1,1,1}
	s := 3
	fmt.Println(findTargetSumWays(nums, s))
}

func findTargetSumWays(nums []int, S int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if len(m) == 0 {
			m[nums[i]] = 1
			m[-1 * nums[i]] = 1
			continue
		}
		t := make(map[int]int)
		for k, v := range m {
			t[k + nums[i]] += v
			t[k - nums[i]] += v
		}
		m = t
	}
	return m[S]
}
