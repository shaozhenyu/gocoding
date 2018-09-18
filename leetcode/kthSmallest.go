package main

import (
	"container/heap"
	"fmt"
)

func main() {
	m := [][]int{[]int{1, 2, 3}, []int{2, 3, 4}}
	fmt.Println(kthSmallest(m, 2))
}

type T struct {
	val  int
	x, y int
}

type THeap []T

func (h THeap) Len() int { return len(h)}

func (h THeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h THeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *THeap) Push(x interface{}) {
	*h = append(*h, x.(T))
}

func (h *THeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	theap := THeap{}
	heap.Init(&theap)
	for j := 0; j < len(matrix[0]); j++ {
		heap.Push(&theap, T{matrix[0][j], 0, j})
	}
	count := 0
	for count < k {
		e := heap.Pop(&theap)
		t := e.(T)
		if t.y < len(matrix[0]) - 1 {
			heap.Push(&theap, T{matrix[t.x][t.y+1], t.x, t.y+1})
		}
		count++
	}
	e := heap.Pop(&theap)
	t := e.(T)
	return t.val
}
