package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 3, 2, 5, 2, 2, 2, 2}
	fmt.Println(majorityElement(a))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	
	max := 0
	major := 0
	before := 0
	numbers := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max = 1
			major = nums[i]
			before = nums[i]
			numbers = 1
		} else if before == nums[i] {
			numbers += 1
			if numbers > max {
				max = numbers
				major = nums[i]
			}
		} else {
			numbers = 1
			before = nums[i]
		}
	}
	return major
}
