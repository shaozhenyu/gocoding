package main

import "fmt"

func main() {
	nums := []int{1,3,2,2,5,2,3,7}
	fmt.Println(findLHS(nums))
}

func findLHS(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	ret := 0
	for k, v := range m {
		if m[k+1] > 0 {
			ret = max(ret, m[k+1] + v)
		}
		if m[k-1] > 0 {
			ret = max(ret, m[k-1] + v)
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
