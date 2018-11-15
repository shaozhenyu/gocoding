package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []Interval{Interval{1, 4}, Interval{2, 3}, Interval{3, 4}}
	fmt.Println(findRightInterval(a))
}

type Interval struct {
	Start int
	End   int
}

type Node struct {
	Start int
	End int
	Idx int
}

func findRightInterval(intervals []Interval) []int {
	s := make([]Node, len(intervals))
	for i := 0; i < len(intervals); i++ {
		s[i] = Node{intervals[i].Start, intervals[i].End, i}
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i].Start == s[j].Start {
			return s[i].End < s[j].End
		}
		return s[i].Start < s[j].Start
	})
	ret := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ret[s[i].Idx] = -1
		if s[i].End > s[len(s)-1].Start {
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[j].Start >= s[i].End {
				ret[s[i].Idx] = s[j].Idx
				break
			}
		}
	}
	return ret
}
