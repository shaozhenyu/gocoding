package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(oddEvenJumps([]int{10, 13, 12, 14, 15}))
	fmt.Println(oddEvenJumps([]int{2, 3, 1, 1, 4}))
	fmt.Println(oddEvenJumps([]int{5, 1, 3, 4, 2}))
	fmt.Println(oddEvenJumps([]int{1, 2, 3, 2, 1, 4, 4, 5}))
}

type Node struct {
	val int
	idx int
}

func oddEvenJumps(A []int) int {
	n := make([]Node, len(A))
	for i := 0; i < len(A); i++ {
		n[i] = Node{A[i], i}
	}
	sort.Slice(n, func(i, j int) bool {
		if n[i].val == n[j].val {
			return n[i].idx < n[j].idx
		}
		return n[i].val < n[j].val
	})
	ret := 0
	odd := make(map[int]bool)
	even := make(map[int]bool)
	for i := 0; i < len(n); i++ {
		if success, ok := odd[i]; ok {
			if success {
				ret++
			}
			continue
		}
		if jump(i, 1, n, odd, even) {
			ret++
		}
	}
	return ret
}

func jump(idx, flag int, n []Node, odd, even map[int]bool) bool {
	if idx >= len(n) {
		return false
	}
	if n[idx].idx == len(n)-1 {
		return true
	}
	if flag%2 == 1 {
		if ret, ok := odd[idx]; ok {
			return ret
		}
		for i := idx + 1; i < len(n); i++ {
			if n[i].idx < n[idx].idx {
				continue
			}
			odd[idx] = jump(i, flag+1, n, odd, even)
			return odd[idx]
		}
		odd[idx] = false
		return odd[idx]
	} else {
		if ret, ok := even[idx]; ok {
			return ret
		}
		for i := idx + 1; i < len(n); i++ {
			if n[i].val == n[idx].val {
				even[idx] = jump(i, flag+1, n, odd, even)
				return even[idx]
			}
		}
		for i := idx - 1; i >= 0; i-- {
			if n[i].idx < n[idx].idx {
				continue
			}
			for j := i - 1; j >= 0; j-- {
				if n[j].idx < n[idx].idx {
					continue
				}
				if n[j].val != n[i].val {
					break
				}
				i = j
			}
			even[idx] = jump(i, flag+1, n, odd, even)
			return even[idx]
		}
		even[idx] = false
		return even[idx]
	}
	return false
}
