package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2,2,3,7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	s := make([]int, 10000)
	sort.Ints(candidates)
	can := []int{candidates[0]}
	for i := 1; i < len(candidates); i++ {
		if candidates[i] != candidates[i-1] {
			can = append(can, candidates[i])
		}
	}
	return fn(0, 0, s, can, target, [][]int{})
}

func fn(start, index int, s []int, candidates []int, target int, ret [][]int) [][]int {
	if target == 0 {
		tmp := make([]int, index)
		for i := 0; i < index; i++ {
			tmp[i] = s[i]
		}
		ret = append(ret, tmp)
		return ret
	}
	if target < 0 {
		return ret
	}
	for i := start; i < len(candidates); i++ {
		s[index] = candidates[i]
		ret = fn(i, index+1, s, candidates, target-s[index], ret)
		s[index] = 0
	}
	return ret
}
