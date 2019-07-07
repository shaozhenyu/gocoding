package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1,-2,-3, 3,-2}))
	fmt.Println(maxSubarraySumCircular([]int{7,9,-8,10,-5}))
	fmt.Println(maxSubarraySumCircular([]int{-1, -2}))
}

// 首先考虑这不是环形数组，那么可以用动态规划很容易的计算出子数组的最大和
// maxNow = Max(A[i], maxNow + A[i]) // 计算直到A[i]的最大值
// max = Max(max, maxNow)
//  eg: 5 -4 3, i=0, maxNow = 5, max = 5
// 				i=1, maxNow = 1, max = 5
// 				i=2, maxNow = 4, max = 5

// 再考虑环形数组 最大值有可能是去掉中间一段负数，那么就可以转换为计算非环形数组的最小值。
// 那么可能的最大值就是 total - min
// minNow = Min(A[i], minNow + A[i])
// min = Min(min, minNow)
// max1 = total - min

// 特殊考虑数组最大值 < 0的情况

func maxSubarraySumCircular(A []int) int {
	max, maxNow := -2 << 31, -2 << 31
	min, minNow := 2 << 31, 2 << 31
	maxVal := -2 << 31
	total := 0
	for i := 0; i < len(A); i++ {
		total += A[i]
		maxVal = Max(maxVal, A[i])

		maxNow = Max(A[i], maxNow + A[i])
		max = Max(max, maxNow)

		minNow = Min(A[i], minNow + A[i])
		min = Min(min, minNow)
	}
	if maxVal <= 0 {
		return maxVal
	}

	return Max(max, total - min)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
