package main

import "fmt"

func main() {
	s := []int{1,20,5}
	fmt.Println(PredictTheWinner(s))
}

func PredictTheWinner(nums []int) bool {
	return predictTheWinner(nums, 0, len(nums)-1) >= 0
}

func predictTheWinner(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	return max(nums[start] - predictTheWinner(nums, start+1, end), nums[end] - predictTheWinner(nums, start, end-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
