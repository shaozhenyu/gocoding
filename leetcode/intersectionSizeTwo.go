package main

import "fmt"

func main() {

}

func intersectionSizeTwo(intervals [][]int) int {
	if len(intervals) <= 1 {
		return len(intervals) * 2
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	s1, s2 := intervals[0][1]-1, intervals[0][1]
	size := 2
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= s1 {
			continue
		}
		if intervals[i][0] > s2 {
			s2 = intervals[i][1] - 1
			size++
		}
        s1 = s2
		s2 = intervals[i][1]
		size++
	}
	return size
}
