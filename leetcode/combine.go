package main

import "fmt"

func main() {
	fmt.Println(combine(5, 3))
}

func combine(n int, k int) [][]int {
	s := make([]int, k)
	return backtrack(0, 1, n, k, s, [][]int{})
}

func backtrack(index, pos, n, k int, s []int, ret [][]int) [][]int {
	if index == k {
		tmp := make([]int, k)
		copy(tmp, s)
		ret = append(ret, tmp)
		return ret
	}
	for i := pos; i <= n; i++ {
		s[index] = i
		ret = backtrack(index+1, i+1, n, k, s, ret)
		s[index] = 0
	}
	return ret
}
