package main

import "fmt"

func main() {
	nums := []int{2, 3, -2, 4, -1, 1, -2, 5}
	nums = []int{-4,-3, 2, -3}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first, last := -1, -1 // idx which < 0
	count := 0
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	max := nums[0]
	if nums[0] < 0 {
		first = 0
		count++
	}
	if nums[0] == 0 {
		return Max(0, maxProduct(nums[1:]))
	}
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			break
		}
		sum[i] = sum[i-1] * nums[i]
		if nums[i] < 0 {
			count++
			if first == -1 {
				first = i
			} else {
				last = i
			}
		}
	}
	if count%2 == 0 {
		max = sum[i-1]
	} else {
		if first - 1 >= 0 {
			max = Max(sum[first-1], max)
		}
		if first < i - 1 && i - 1 >= 0 {
			max = Max(sum[i-1]/sum[first], max)
		}
		if last - 1 >= 0 {
			max = Max(sum[last-1], max)
		}
		if last >= 0 && last < i - 1 && i - 1 >= 0 {
			max = Max(sum[i-1]/sum[last], max)
		}
	}
	if i < len(nums) {
		max = Max(max, Max(0, maxProduct(nums[i+1:])))
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
