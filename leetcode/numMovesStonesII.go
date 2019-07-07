package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStonesII([]int{7,4,9}), []int{1,2})
	fmt.Println(numMovesStonesII([]int{100,101,104,102,103}), []int{0, 0})
	fmt.Println(numMovesStonesII([]int{6,5,4,3,10}), []int{2, 3})
	fmt.Println(numMovesStonesII([]int{8,7,6,5,10}), []int{1, 1})
	fmt.Println(numMovesStonesII([]int{1,500000000,1000000000}), []int{2,499999999})
}

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	max := 0
	min := 2 << 31
	for i := 1; i < len(stones); i++ {
		max += (stones[i] - stones[i-1] - 1)
	}
	if max == 0 {
		return []int{0, 0}
	}
	if max - (stones[1] - stones[0] - 1) == 0 || max - (stones[len(stones)-1] - stones[len(stones)-2] - 1) == 0 {
		if max - (stones[1] - stones[0] - 1) == 0 {
			min = Min(2, stones[1] - stones[0] - 1)
		} else {
			min = Min(2, stones[len(stones)-1] - stones[len(stones)-2] - 1)
		}
		return []int{min, max}
	}
	max -= Min(stones[1] - stones[0] - 1, stones[len(stones)-1] - stones[len(stones)-2] - 1)
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[j] - stones[i] + 1 >= len(stones) {
				min = Min(min, stones[j] - stones[i] - 1 - (j - i - 1))
				break
			} else {
				min = Min(min, len(stones) - (j - i + 1))
			}
		}
	}
	min = Min(min, len(stones) - 1)
	return []int{min, max}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
