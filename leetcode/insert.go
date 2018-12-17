package main

import "fmt"

func main() {
	intervals := []Interval{Interval{1,2},Interval{3,5},Interval{6,7},Interval{8,10},Interval{12,16}}
	newInterval := Interval{4, 8}
	fmt.Println(insert(intervals, newInterval))
}

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	ret := []Interval{}
	for i := 0; i < len(intervals); i++ {
		if newInterval.End < intervals[i].Start {
			t := append([]Interval{}, intervals[i:]...)
			ret = append(intervals[:i], newInterval)
			ret = append(ret, t...)
			return ret
		}
		if newInterval.End <= intervals[i].End {
			if newInterval.Start <= intervals[i].Start {
				intervals[i].Start = newInterval.Start
			}
			return intervals
		}
		if newInterval.Start <= intervals[i].End && newInterval.End > intervals[i].End {
			if newInterval.Start > intervals[i].Start {
				newInterval.Start = intervals[i].Start
			}
			intervals = append(intervals[:i],intervals[i+1:]...)
			i--
		}
	}
	intervals = append(intervals, newInterval)
	return intervals
}
