package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func main() {
	i := []Interval{Interval{1, 2}, Interval{2, 3}, Interval{3, 4}, Interval{1, 3}}
	fmt.Println(eraseOverlapIntervals(i))
}

func eraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})
	count := 0
	end := intervals[0].End
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start >= end {
			end = intervals[i].End
		} else {
			if intervals[i].End < end {
				end = intervals[i].End
			}
			count++
		}
	}
	return count
}
