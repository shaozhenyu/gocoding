package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(lastStoneWeight([]int{4,3,3,1,1}))
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, s := range stones {
		heap.Push(h, s)
	}
	for h.Len() > 1 {
		x1 := heap.Pop(h).(int)
		x2 := heap.Pop(h).(int)
		if x1 != x2 {
			heap.Push(h, x1 - x2)
		}
	}
	if h.Len() == 1 {
		return heap.Pop(h).(int)
	}
	return 0
}
