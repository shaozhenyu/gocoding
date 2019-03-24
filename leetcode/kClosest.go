package main

import (
	"fmt"
	"container/heap"
)

func main() {
	points := [][]int{[]int{1,3}, []int{2, -2}, []int{10,1}}
	K := 2
	fmt.Println(kClosest(points, K))
}

type Node struct {
	d int
	p []int
}

type NHeap []Node

func (h NHeap) Len() int           { return len(h) }
func (h NHeap) Less(i, j int) bool { return h[i].d > h[j].d }
func (h NHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, K int) [][]int {
	h := &NHeap{}
	heap.Init(h)
	for i := 0; i < len(points); i++ {
		n := Node{points[i][0] * points[i][0] + points[i][1] * points[i][1], points[i]}
		heap.Push(h, n)
		if h.Len() > K {
			heap.Pop(h)
		}
	}
	ret := [][]int{}
	for h.Len() > 0 {
		r := heap.Pop(h).(Node)
		ret = append(ret, r.p)
	}
	return ret
}

