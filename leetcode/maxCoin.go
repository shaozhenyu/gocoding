package main

import "fmt"

func main() {
	nums := []int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3}
	fmt.Println(maxCoins(nums))
}

func maxCoins(nums []int) int {
	max := 0
	for i := 1; i < len(nums) - 1; i++ {
		this := coins(nums[:i]) * nums[i] * coins(nums[i+1:])
		fmt.Println(this)
		if this > max {
			max = this
		}
	}
	return max
}

func coins(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0
	for i := 0; i < len(nums) - 1; i++ {
		this := coins(nums[:i]) * nums[i] * coins(nums[i+1:])
		if this > max {
			max = this
		}
	}
	return max
}
