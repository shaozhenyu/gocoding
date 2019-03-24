package main

import (
	"fmt"
	"container/heap"
)

func main() {
	kt := Constructor(2, []int{4,2,1,6,7})
	fmt.Println(kt.Add(3))
	fmt.Println(kt.Add(10))
	fmt.Println(kt.Add(11))
	fmt.Println(kt.Add(8))
}

type KthLargest struct {
	k int
	min int
	h *IntHeap
}


func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	heap.Init(&h)
	for h.Len() > k {
		heap.Pop(&h)
	}
	return KthLargest{k, &h}
}


func (this *KthLargest) Add(val int) int {
	if val <= min && this.h.Len() >= this.k {
		return min
	}
	heap.Push(this.h, val)
	for this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	t := heap.Pop(this.h).(int)
	heap.Push(this.h, t)
	return t
}


type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
