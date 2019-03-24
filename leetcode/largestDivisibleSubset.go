package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,1,1,2}
	fmt.Println(largestDivisibleSubset(nums))
}

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	dp[len(nums)-1] = []int{nums[len(nums)-1]}
	ret := dp[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		t := make([]int, 0)
		for j := i + 1; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 && len(dp[j]) > len(t) {
				t = dp[j]
			}
			if len(t) >= len(ret) {
				break
			}
		}
		dp[i] = append([]int{nums[i]}, t...)
		if len(dp[i]) > len(ret) {
			ret = dp[i]
		}
	}
	return ret
}
