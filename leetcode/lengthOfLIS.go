package main

import "fmt"

func main() {
	nums := []int{4, 10, 4, 3, 8, 9}
	fmt.Println(lengthOfLIS(nums))
}

// O(NlogN)
func lengthOfLIS(nums []int) int {
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
	dp := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = 1
		max := 0
		for j := i; j < len(nums); j++ {
			if nums[i] < nums[j] {
				if dp[j] > max {
					max = dp[j]
				}
			}
		}
		dp[i] += max
	}
	ret := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	fmt.Println(dp)
	return ret
}
