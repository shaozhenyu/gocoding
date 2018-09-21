package main

import "fmt"

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1, 13}))
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max := -1 << 32
	min := 1 << 32
	for i := 0; i < len(nums); i++ {
		max = Max(nums[i], max)
		min = Min(nums[i], min)
	}
	maxB := make([]int, len(nums)+1)
	minB := make([]int, len(nums)+1)
	hasB := make([]bool, len(nums)+1)
	size := len(nums)
	for i := 0; i < len(nums); i++ {
		index := (nums[i] - min) * size / (max - min)
		if hasB[index] {
			maxB[index] = Max(nums[i], maxB[index])
			minB[index] = Min(nums[i], minB[index])
		} else {
			hasB[index] = true
			maxB[index] = nums[i]
			minB[index] = nums[i]
		}
	}

	pre := maxB[0]
	ret := -1 << 32
	for i := 1; i <= len(nums); i++ {
		if hasB[i] {
			ret = Max(ret, minB[i]-pre)
			pre = maxB[i]
		}
	}
	return ret
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
