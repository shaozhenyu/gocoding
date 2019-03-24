package main

import "fmt"

func main() {
	nums := []int{10,5,2,6}
	k := 100
	fmt.Println(numSubarrayProductLessThanK(nums, k))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	start := 0
	end := 0
	count := 0
	for start < len(nums) {
		if sum < k && end < len(nums) {
			count += end - start + 1
			end++
			if end < len(nums) {
				sum *= nums[end]
			}
		} else {
			sum /= nums[start]
			start++
			if start > end {
				end = start
			}
		}
	}
	return count
}
