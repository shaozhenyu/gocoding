package main

import "fmt"

func main() {
	nums := []int{4, 5, 5, 5, 5, 0, 1, 2, 3}
	fmt.Println(0, findMin(nums))

	nums = []int{2, 2, 2, 0, 1}
	fmt.Println(0, findMin(nums))

	nums = []int{2, 2, 0, 1, 1, 1, 1, 1, 1}
	fmt.Println(0, findMin(nums))

	nums = []int{3, 1, 3}
	fmt.Println(1, findMin(nums))

	nums = []int{1,1,1,1,2,2,2,2,2}
	fmt.Println(1, findMin(nums))

	nums = []int{2,1,1,1}
	fmt.Println(1, findMin(nums))

	fmt.Println("aaaaaa")
	nums = []int{3,3,1,3}
	fmt.Println(1, findMin(nums))
}

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	for nums[low] >= nums[high] {
		mid := (low + high) / 2
		if nums[low] == nums[mid] && nums[high] == nums[mid] {
			break
		}
		fmt.Println(nums[low], nums[high], nums[mid])
		if nums[mid] >= nums[low] {
			low = mid + 1
		} else if nums[mid] <= nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = high
		}
	}
	return nums[low]
}
