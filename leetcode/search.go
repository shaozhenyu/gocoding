package main

import "fmt"

func main() {
	v := []int{1, 3, 5}
	fmt.Println(search(v, 3))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums) - 1
	ok := false
	for nums[start] >= nums[end] {
		ok = true
		mid := (start + end)/2
		fmt.Println(start, end, mid)
		if nums[mid] > nums[start] {
			start = mid
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			break
		}
	}
	fmt.Println("aa:", start, end)
	if !ok {
		return binarySearch(nums, target)
	}
	if nums[start] >= target && nums[0] <= target {
		return binarySearch(nums[:start+1], target)
	}
	pos := binarySearch(nums[end:], target)
	if pos == -1 {
		return -1
	}
	return end + pos
}

func binarySearch(nums []int, target int) int {
	start, end := 0, len(nums) - 1
	for start <= end {
		fmt.Println(nums, start, end)
		mid := (start + end)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}