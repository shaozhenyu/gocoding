package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	fmt.Println(canPartitionKSubsets(nums, k))
}

func canPartitionKSubsets(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	if len(nums) < k {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	avg := sum/k
	sort.Ints(nums)
	fmt.Println(nums)
	if nums[len(nums)-1] > avg {
		return false
	}
	visited := make([]bool, len(nums))
	for i := 0; i < k; i++ {
		if !backtrace(avg, nums, visited) {
			return false
		}
	}
	return true
}

func backtrace(target int, nums []int, visited []bool) bool {
	if target == 0 {
		return true
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if visited[i] {
			continue
		}
		if nums[i] > target {
			continue
		}
		visited[i] = true
		if backtrace(target-nums[i], nums, visited) {
			return true
		}
		visited[i] = false
	}
	return false
}


