package main

import (
	"fmt"
)

func main() {
	num := []int{1,5,5,11,12}
	fmt.Println(canPartition(num))
}

func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum % 2 != 0 {
		return false
	}
	half := sum / 2
	half -= nums[0]
	if half < 0 {
		return false
	} else if half == 0 {
		return true
	}

	nums = nums[1:]
	dp := make([]bool, half+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := half; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j - nums[i]]
		}
		if dp[half] == true {
			return true
		}
	}
	return dp[half]
}
