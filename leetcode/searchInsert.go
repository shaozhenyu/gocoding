package main

import "fmt"

func main() {
	nums := []int{1,3,5,6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mid := (lo + hi)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}
