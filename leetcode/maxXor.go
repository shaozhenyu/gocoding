package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaximumXOR([]int{8,2,10}))
}

func findMaximumXOR(nums []int) int {
	lenN := len(nums)
	maxN := nums[0]
	for i := 1; i < lenN; i++ {
		if nums[i] > maxN {
			maxN = nums[i]
		}
	}
	fmt.Println(maxN)
	max := 0
	for i := 0; i < lenN; i++ {
		tmp := maxN ^ nums[i]
		fmt.Println(tmp)
		if tmp > max {
			max = tmp
		}
	}
	
	return max
}
