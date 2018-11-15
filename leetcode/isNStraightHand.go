package main

import (
	"fmt"
	"sort"
)

func main() {
	hand := []int{1,2,3,2,3,4,6,7,8}
	W := 3
	fmt.Println(isNStraightHand(hand, W))
}

type Node struct {
	v int
	count int
}

func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}
	m := make(map[int]int)
	for i := 0; i < len(hand); i++ {
		m[hand[i]]++
	}
	s := make([]Node, len(m))
	idx := 0
	for k, v := range m {
		s[idx] = Node{k, v}
		idx++
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].v < s[j].v
	})
	fmt.Println(s)
	for len(s) > 0 {
		d := 1
		pre := s[0].v
		count := s[0].count
		for i := 1; i < W; i++ {
			if i >= len(s) {
				return false
			}
			if s[i].v != pre + 1 || s[i].count < count {
				return false
			}
			pre = s[i].v
			s[i].count -= count
			if s[i].count == 0 {
				d++
			}
		}
		s = s[d:]
	}
	return true
}
