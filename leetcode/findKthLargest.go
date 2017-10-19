package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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
	nums := []int{1, 4, 12, 2, 4, 1, 2, 4, 3, 3, 1, 3, 4, 3, 5, 1}
	findKthLargest(nums, 5)
}

func findKthLargest(nums []int, k int) int {
	hh := IntHeap(nums)
	h := &hh

	heap.Init(h)
	for k > 0 {
		fmt.Println(heap.Pop(h))
		k--
	}
	return 0
}
