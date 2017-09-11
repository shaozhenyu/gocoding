package main

import (
	"fmt"
	"sort"
)

func main() {
	//[[1 6] [2 8] [7 12] [10 16]]
	a := [][]int{[]int{1, 6}, []int{2, 8}, []int{7, 12}, []int{10, 16}}
	fmt.Println(findMinArrowShots(a))
}

type NeedSort [][]int

func (n NeedSort) Len() int {
	return len(n)
}

func (n NeedSort) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NeedSort) Less(i, j int) bool {
	if n[i][0] != n[j][0] {
		return n[i][0] < n[j][0]
	} else {
		return n[i][1] < n[j][1]
	}
	return false
}

func findMinArrowShots(points [][]int) int {
	if len(point) == 0 {
		return 0
	}
	p := NeedSort(points)
	sort.Sort(p)
	fmt.Println(p)
	min := 1
	for i := 1; i < len(p); i++ {
		if p[i-1][1] < p[i][0] {
			min++
		} else {
			if p[i-1][1] < p[i][1] {
				p[i][1] = p[i-1][1]
			}
		}
	}
	return min
}
