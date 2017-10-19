package main

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{1,2,2,4}))
}

func findErrorNums(nums []int) []int {
	m := make(map[int]int, len(nums))

	for i := 1; i <= len(nums); i++ {
		m[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			delete(m, nums[i])
		} else {
			m[nums[i]] = 1
		}
		fmt.Println(m)
	}

	fmt.Println(m)
	ret := make([]int, 2)
	i := 0
	for k, _ := range m {
		ret[i] = k
		i++
	}
	return ret
}
