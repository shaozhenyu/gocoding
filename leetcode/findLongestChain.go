package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestChain([][]int{[]int{-6, 9}, []int{1, 6}, []int{8, 10}, []int{-1, 4}, []int{-6, -2}, []int{-9, 8}, []int{-5, 3}, []int{0, 3}}))
}

type NeedSort [][]int

func (n NeedSort) Len() int {
	return len(n)
}

func (n NeedSort) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NeedSort) Less(i, j int) bool {
	if n[i][1] != n[j][1] {
		return n[i][1] < n[j][1]
	} else {
		return n[i][0] < n[j][0]
	}
	return false
}

func findLongestChain(pairs [][]int) int {
	p := NeedSort(pairs)
	sort.Sort(p)

	max := 1

	fmt.Println(p)

	lenP := len(p)
	now := 0
	for i := 1; i < lenP; i++ {
		if p[i][1] == p[now][1] || p[i][0] <= p[now][1] {
			continue
		}
		max++
		now = i
	}

	return max
}
