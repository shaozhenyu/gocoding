package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rearrangeBarcodes([]int{1,1,1,2,2,2}))
}

func rearrangeBarcodes(barcodes []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(barcodes); i++ {
		m[barcodes[i]]++
	}
	n := make([]Node, 0, len(m))
	for k, v := range m {
		n = append(n, Node{k, v})
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i].count > n[j].count
	})
	fmt.Println(n)
	ret := make([]int, len(barcodes))
	start := 0
	for _, v := range n {
		val, count := v.val, v.count
		for count > 0 {
			fmt.Println(start)
			if ret[start] == 0 {
				ret[start] = val
				start += 2
				count--
			} else {
				start++
			}
			start = start%len(ret)
		}
	}
	return ret
}

type Node struct {
	val int
	count int
}
