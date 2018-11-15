package main

import "fmt"

func main() {
	s := 7
	nums := []int{1,2,3,4,5,1,2,6,7}
	s = 15
	nums = []int{1,2,3,4,5}
	fmt.Println(minSubArrayLen(s, nums))
}

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, 0
	sum := 0
	ret := len(nums) + 1
	for start < len(nums) && end < len(nums) {
		sum += nums[end]
		for sum >= s {
			if end - start + 1 < ret {
				ret = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		if ret == 1 {
			return ret
		}
		end++
	}
	if ret == len(nums) + 1 {
		return 0
	}
	return ret
}
