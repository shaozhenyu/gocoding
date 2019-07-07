package main

import "fmt"

func main() {
	fmt.Println(clumsy(10))
}

func clumsy(N int) int {
	ret := 0
	nums := make([]int, 0)
	for N > 0 {
		t := N
		N--
		if N >= 1 {
			t *= N
		}
		N--
		if N >= 1 {
			t /= N
		}
		N--
		if N >= 1 {
			ret += N
		}
		N--
		nums = append(nums, t)
	}
	fmt.Println(nums)
	ret += nums[0]
	for i := 1; i < len(nums); i++ {
		ret -= nums[i]
	}
	return ret
}
