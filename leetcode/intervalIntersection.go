package main

import "fmt"

func main() {
	a := []Interval{Interval{0, 1}}
	b := []Interval{Interval{0, 1}}
	fmt.Println(intervalIntersection(a, b))
}


type Interval struct {
	Start int
	End   int
}

func intervalIntersection(A []Interval, B []Interval) []Interval {
	ret := []Interval{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if B[j].End < A[i].Start {
			j++
			continue
		}
		if B[j].End <= A[i].End {
			start := max(A[i].Start, B[i].Start)
			ret = append(ret, Interval{start, B[j].End})
			j++
			continue
		}
		if B[j].Start >= A[i].Start && B[j].Start <= A[i].End {
			end := min(A[i].End, B[j].End)
			ret = append(ret, Interval{B[j].Start, end})
			if A[i].End < B[j].End {
				i++
			} else {
				j++
			}
			continue
		}
		if B[j].Start > A[i].End {
			i++
			continue
		}
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
