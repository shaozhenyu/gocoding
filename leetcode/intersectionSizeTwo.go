package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{[]int{1, 3}, []int{1, 4}, []int{2, 5}, []int{3, 5}}
	fmt.Println(intersectionSizeTwo(intervals), 3)
	intervals = [][]int{[]int{1, 2}, []int{2, 3}, []int{2, 4}, []int{4, 5}}
	fmt.Println(intersectionSizeTwo(intervals), 5)
	intervals = [][]int{[]int{1, 3}, []int{4, 9}, []int{0, 10}, []int{6, 7}, []int{1, 2}, []int{0, 6}, []int{7, 9}, []int{0, 1}, []int{2, 5}, []int{6, 8}}
	fmt.Println(intersectionSizeTwo(intervals), 7)
	intervals = [][]int{[]int{4, 7}, []int{5, 8}, []int{7, 9}, []int{2, 6}, []int{0, 1}, []int{1, 4}, []int{1, 9}, []int{0, 5}, []int{5, 10}, []int{7, 8}}
	fmt.Println(intersectionSizeTwo(intervals), 6)
}

func intersectionSizeTwo(intervals [][]int) int {
	if len(intervals) <= 1 {
		return len(intervals) * 2
	}
	m := make(map[int]bool)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		fmt.Println(pre, m)
		if intervals[i][0] <= pre[0] {
			continue
		}
		if intervals[i][0] < pre[1] {
			if m[pre[0]] && pre[1] < intervals[i][1]{
				pre = []int{pre[1], intervals[i][1]}
			} else {
				pre = []int{intervals[i][0], pre[1]}
			}
			continue
		}
		if intervals[i][0] >= pre[1] {
			m[pre[1]] = true
			if !m[pre[0]] && !m[pre[1]-1] {
				m[pre[1]-1] = true
			}
			pre = intervals[i]
		}
		fmt.Println(m)
	}
	m[pre[0]] = true
	m[pre[0]+1] = true
	fmt.Println("last m : ", m)
	return len(m)
}
