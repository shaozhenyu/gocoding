package main

import "fmt"

func main() {
	nums := []int{2,0}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	reach := nums[0]
	start := 1
	for start <= reach && start < len(nums) {
		if start + nums[start] > reach {
			reach = start + nums[start]
		}
		if reach >= len(nums) - 1 {
			return true
		}
		start++
	}
	return false
}
