package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 2, 2, 5, 5, 1, 1, 2, 2, 1, 1, 1, 1}
	nums = []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}
	fmt.Println(makesquare(nums))
}

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}
	max := 0
	sum := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		max = Max(max, nums[i])
		m[nums[i]]++
	}
	avg := sum / 4
	if sum%4 > 0 || max > avg {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		ok := ms(len(nums)-1, avg, nums, m)
		if !ok {
			return false
		}
	}
	for k, v := range m {
		if v <= 0 {
			delete(m, k)
		}
	}
	return len(m) == 0
}

func ms(idx, target int, nums []int, m map[int]int) bool {
	if target == 0 {
		return true
	}
	if len(nums) == 0 {
		return false
	}
	for i := idx; i >= 0; i-- {
		if m[nums[i]] <= 0 {
			continue
		}
		if target < nums[i] {
			continue
		}
		m[nums[i]] -= 1
		if target == nums[i] {
			return true
		}
		ok := ms(i-1, target-nums[i], nums, m)
		if ok {
			return ok
		}
		m[nums[i]] += 1
	}
	return false
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
