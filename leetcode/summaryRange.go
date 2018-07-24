package main

type SummaryRanges struct {
	startM map[int]int
	endM   map[int]int
	st     []int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{
		startM: map[int]int{},
		endM:   map[int]int{},
		st:     []int{},
	}
}

func (this *SummaryRanges) Addnum(val int) {
	startM, endM := this.startM, this.endM
	flag := false
	for start, end := range startM {
		if val >= start && val <= end {
			return
		} else if val == end+1 {
			startM[start] = val
			delete(endM, end)
			endM[val] = start
			if _, ok := startM[val+1]; ok {
				startM[start] = startM[val+1]
				endM[startM[val+1]] = start
				delete(startM, val+1)
				delete(endM, val)
			}
			flag = true
			break
		} else if val == start-1 {
			startM[val] = end
			delete(startM, start)
			endM[end] = val
			if _, ok := endM[val-1]; ok {
				startM[endM[val-1]] = end
				endM[end] = endM[val-1]
				delete(endM, val-1)
				delete(startM, val)
			}
			flag = true
			break
		}
	}
	if !flag {
		startM[val] = val
		endM[val] = val
	}
	s := make([]int, len(startM))
	index := 0
	for key := range startM {
		s[index] = key
		index++
	}
	sort.Ints(s)
	this.st = s
}

func (this *SummaryRanges) Getintervals() []Interval {
	s := this.st
	m := this.startM
	ret := make([]Interval, len(s))
	for i := 0; i < len(s); i++ {
		ret[i] = Interval{s[i], m[s[i]]}
	}
	return ret
}
