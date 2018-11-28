package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{0,0,4,4}
	// citations = []int{11, 15}
	fmt.Println(hIndex(citations))
}

//  “一位有 h 指数的学者，代表他（她）的 N 篇论文中至多有 h 篇论文，分别被引用了至少 h 次，其余的 N - h 篇论文每篇被引用次数不多于 h 次。”
func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		h := len(citations) - i
		if citations[i] >= h {
			return h
		}
	}
	return len(citations)
}
