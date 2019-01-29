package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 2, 1, 2, 6, 7, 5, 1, 6, 3, 8, 11, 4, 3, 2, 6, 7}
	//nums := []int{1, 2, 1, 2, 6, 7, 5, 1}
	nums := []int{17, 7, 19, 11, 1, 19, 17, 6, 13, 18, 2, 7, 12, 16, 16, 18, 9, 3, 19, 5}
	k := 6
	fmt.Println(maxSumOfThreeSubarrays(nums, k))
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	array := make([]int, len(nums)-k+1)
	sum := 0
	pos := 0
	for i := 0; i < len(nums); i++ {
		if i >= k {
			array[pos] = sum
			pos++
			sum = sum + nums[i] - nums[i-k]
		} else {
			sum += nums[i]
		}
	}
	array[pos] = sum

	left := make([][]int, len(array))
	right := make([][]int, len(array))
	var maxLeft, maxRight int
	var leftPos, rightPos int
	for i, j := k, len(array)-k-1; i < len(array) && j >= 0; i, j = i+1, j-1 {
		left[i] = []int{0, 0}
		right[j] = []int{0, 0}
		if array[i-k] > maxLeft {
			maxLeft = array[i-k]
			leftPos = i - k
		}
		left[i][0] = maxLeft
		left[i][1] = leftPos
		if array[j+k] > maxRight {
			maxRight = array[j+k]
			rightPos = j + k
		}
		right[j][0] = maxRight
		right[j][1] = rightPos
	}

	max := 0
	var a, b, c int
	for i := k; i < len(array)-k; i++ {
		v := array[i] + left[i][0] + right[i][0]
		if v > max {
			max = v
			a = left[i][1]
			b = i
			c = right[i][1]
		}
	}
	return []int{a, b, c}
}
