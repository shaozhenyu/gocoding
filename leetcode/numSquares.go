package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(132))
}

func numSquares(n int) int {
	dp := make([]int, n + 1)

	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = 2 << 31
		for j := 1; j * j <= i; j++ {
			count := dp[i - j*j] + 1
			if count < dp[i] {
				dp[i] = count
			}
		}
	}
	return dp[n]
}

func numSquares1(n int) int {
	sqrt := int(math.Sqrt(float64(n)))
	nums := make([]int, sqrt)
	for i := 1; i <= sqrt; i++ {
		nums[i-1] = i * i
	}
	fmt.Println(nums)

	return backtrack(nums, n, 0, n)
}

func backtrack(nums []int, n int, count int, min int) int {
	if count >= min {
		return min
	}
	this := n
	for i := len(nums)-1; i >= 0; i-- {
		this = this - nums[i]
		if this < 0 {
			this = this + nums[i]
			continue
		}
		if this == 0 {
			if count + 1 < min {
				min = count + 1
			}
			this = this + nums[i]
			continue
		}
		newC := backtrack(nums, this, count+1, min)
		if newC < min {
			min = newC
		}
		this = this + nums[i]
	}
	return min
}
