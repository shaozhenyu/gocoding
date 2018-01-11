package main

import "fmt"

func main() {
	nums := []int{1,12,-5,-6,50,3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	pre := 0
	var now int
	for i := k; i < len(nums); i++ {
		now = max + nums[i] - nums[pre]
		if now > max {
			max = now
		}
		pre++
	}
	return float64(max)/float64(k)
}
