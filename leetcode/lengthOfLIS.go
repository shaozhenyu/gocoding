package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	size := 0
	tail := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if size == 0 || tail[size-1] < nums[i] {
			tail[size] = nums[i]
			size++
		} else {
			x, y := 0, size-1
			for x < y {
				mid := (x + y) / 2
				if tail[mid] < nums[i] {
					x = mid + 1
				} else {
					y = mid
				}
			}
			tail[x] = nums[i]
		}
	}
	return size
}

// O(NlogN)
func lengthOfLIS2(nums []int) int {
	tail := make([]int, len(nums))
	size := 0
	for _, num := range nums {
		if size == 0 || num > tail[size-1] {
			tail[size] = num
			size++
		} else {
			i := 0
			j := size
			for i != j {
				mid := (i + j) / 2
				if tail[mid] < num {
					i = mid + 1
				} else {
					j = mid
				}
			}
			tail[i] = num
		}

		fmt.Println(tail, size)
	}
	return size
}

func lengthOfLIS1(nums []int) int {
	maxEnd := make([]int, len(nums))
	maxEnd[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				maxEnd[i] = max(maxEnd[j]+1, maxEnd[i])
			}
		}
	}
	fmt.Println(maxEnd)
	ret := 0
	for i := 0; i < len(maxEnd); i++ {
		ret = max(maxEnd[i], ret)
	}
	return ret + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
