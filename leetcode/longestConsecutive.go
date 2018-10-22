package main

import "fmt"

func main() {
	nums := []int{200,1,2,10,3,4,100}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = 1
	}
	max := 1
	for k := range m {
		delete(m, k)
		count := 1
		k1 := k - 1
		for m[k1] > 0 {
			delete(m, k1)
			count++
			k1--
		}
		k1 = k + 1
		for m[k1] > 0 {
			delete(m, k1)
			count++
			k1++
		}
		if count > max {
			max = count
		}
	}
	return max
}
