package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestDistancePair([]int{1,1,2,3,4}, 3))
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	low, high := 0, nums[len(nums)-1] - nums[0]
	for low < high {
		mid := low + (high - low)/2
		count := 0
		for i := 1; i < len(nums); i++ {
			j := 0
			for nums[i] - nums[j] > mid {
				j++
			}
			count += i - j
		}
		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
