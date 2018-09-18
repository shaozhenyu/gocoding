package main

import "fmt"

func main() {
	// p := [][]int{[]int{1, 2}, []int{2, 3}, []int{6, 1}, []int{2, 1}, []int{2, 2}, []int{100, 100}, []int{200, 100}}
	p := [][]int{[]int{1,2},[]int{1,3}}
	fmt.Println(fallingSquares(p))
}

type LinkList struct {
	start int
	end   int
	high  int
	next  *LinkList
}

func fallingSquares(positions [][]int) []int {
	list := &LinkList{}
	ret := make([]int, len(positions))
	high := 0
	for i := 0; i < len(positions); i++ {
		v := positions[i]
		start := v[0] + 1
		end := start + v[1]
		max := list.getMaxInPos(start, end)
		fmt.Println("max: ", max, start, end, i)
		now := max + v[1]
		list.insert(start, end, now)
		fmt.Println("insert ok")
		if high < now {
			high = now
		}
		ret[i] = high
	}
	return ret
}

func (l *LinkList) insert(start, end, high int) {
	newL := &LinkList{
		start: start,
		end:   end,
		high:  high,
	}
	if l.next == nil {
		l.next = newL
		return
	}
	t := l.next
	if start <= t.start {
		newL.next = t
		l.next = newL
		return
	}
	pre := t
	t = t.next
	for t != nil && start > t.start {
		pre = t
		t = t.next
	}
	if t == nil {
		pre.next = newL
	} else {
		newL.next = t
		pre.next = newL
	}
}

func (l *LinkList) getMaxInPos(start, end int) (high int) {
	t := l.next
	for t != nil {
		if (start > t.start && start < t.end) || (end > t.start && end < t.end) || (start <= t.start && end >= t.end){
			if high < t.high {
				high = t.high
			}
		}
		// if start >= t.end {
		// 	break
		// }
		t = t.next
	}
	return
}
