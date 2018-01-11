package main

import "fmt"

func main() {
	nums := []int{7, 2, 5, 10, 8}
	m := 2
	fmt.Println(splitArray(nums, m))
}

func splitArray(nums []int, m int) int {
	max, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		sum += nums[i]
	}

	valid := func(max, m int, nums []int) bool {
		sum := 0
		count := 0
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
			if sum > max {
				sum = nums[i]
				count++
				if count >= m {
					return false
				}
			}
		}
		return true
	}

	left, right := max, sum

	var mid int
	for left <= right {
		fmt.Println(left, right)
		mid = (left + right) / 2
		if valid(mid, m, nums) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
