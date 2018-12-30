package main

import "fmt"

func main() {
	nums := []int{1,2,2,5,3,5}
	nums = []int{1,-2147483648,2}
	nums = []int{-2147483648,-2147483648,-2147483648,-2147483648,1,1,1}
	fmt.Println(thirdMax(nums))
}

func thirdMax(nums []int) int {
	a, b, c := -1 << 31, -1 << 31, -1 << 31
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
		if nums[i] > a {
			a, b, c = nums[i], a, b
		} else if nums[i] < a && nums[i] > b {
			b, c = nums[i], b
		} else if nums[i] < b && nums[i] > c {
			c = nums[i]
		}
	}
	fmt.Println(a, b, c, m)
	if len(m) < 3 {
		return a
	}
	return c
}
