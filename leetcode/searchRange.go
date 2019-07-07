package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 2, 3, 3, 3, 3, 4, 4, 4}
	target := 2
	fmt.Println(searchRange(nums, target))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
}

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums) - 1
	for left < right {
		mid := left + (right - left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if nums[left] == target && nums[right] == target {
				break
			}
			if nums[left] < target {
				left++
			}
			if nums[right] > target {
				right--
			}
		}
	}
	return []int{left, right}
}

func searchRange1(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	lo, li := 0, len(nums)-1
	var mid int
	for lo <= li {
		mid = (lo + li) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			break
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			li = mid - 1
		}
	}
	ret := []int{-1, -1}
	fmt.Println(mid, nums[mid])
	if nums[mid] == target {
		i := mid - 1
		for ; i >= 0 && nums[i] == target; i-- {
		}
		ret[0] = i + 1
		i = mid + 1
		for ; i < len(nums) && nums[i] == target; i++ {
		}
		ret[1] = i - 1
	}
	return ret
}
