package main

import "fmt"

func main() {
	nums := []int{9,2,1,4,5,3,7}
	nums = []int{1,1}
	fmt.Println("old:", nums)
	nextPermutation(nums)
	fmt.Println("new:", nums)
}

func nextPermutation(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			idx := i+1
			for j := i + 2; j < len(nums); j++ {
				if nums[j] < nums[idx] && nums[j] > nums[i] {
					idx = j
				}
			}
			nums[i], nums[idx] = nums[idx], nums[i]
			quicksort(nums, i+1, len(nums)-1)
			return
		}
	}
	quicksort(nums, 0, len(nums) - 1)
}

func quicksort(nums []int, low, high int) {
	if low >= high {
		return
	}
	mid := nums[low]
	i, j := low, high
	for i < j {
		fmt.Println("xx,", i, j)
		for i < j && nums[j] >= mid {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= mid {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = mid
	quicksort(nums, low, i-1)
	quicksort(nums, j+1, high)
}
