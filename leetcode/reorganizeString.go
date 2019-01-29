package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "baaabcc"
	fmt.Println(reorganizeString(s))
}

type Node struct {
	val byte
	count int
}

func reorganizeString(S string) string {
	m := make(map[byte]int)
	max := 0
	for i := 0; i < len(S); i++ {
		m[S[i]]++
		if m[S[i]] > max {
			max = m[S[i]]
		}
	}
	t := len(S)/2
	if len(S)%2 > 0 {
		t += 1
	}
	if max > t {
		return ""
	}
	n := make([]Node, 0, len(m))
	for k, v := range m {
		n = append(n, Node{k, v})
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i].count > n[j].count
	})
	idx := 0
	ret := make([]byte, len(S))
	for i := 0; i < len(n); i++ {
		b := n[i].val
		count := n[i].count
		for count > 0 {
			ret[idx] = b
			idx += 2
			if idx >= len(S) {
				idx = 1
			}
			count--
		}
	}
	return string(ret)
}
