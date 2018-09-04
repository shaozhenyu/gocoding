package main

import "fmt"

func main() {
	fmt.Println(findMaximumXOR([]int{2, 8, 10}))
}

func findMaximumXOR(nums []int) int {
	max := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask |= 1 << uint(i)
		m := make(map[int]bool)
		for _, n := range nums {
			m[n & mask] = true
		}
		tmp := max | 1 << uint(i)
		for k := range m {
			if _, ok := m[tmp ^ k]; ok {
				max = tmp
				break
			}
		}
	}
	return max
}
