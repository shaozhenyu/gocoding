package main

import (
	"fmt"
	"container/heap"
)

func main() {
	A := []int{1, 2, 3, 5}
	K := 3
	A = []int{1,2,11,37,83,89}
	K = 11
	fmt.Println(A)
	fmt.Println(kthSmallestPrimeFraction(A, K))
}

type Node struct {
	x, y int
	val float64
}

func kthSmallestPrimeFraction(A []int, K int) []int {
	h := &NHeap{}
	heap.Init(h)
	for i := 0; i < len(A); i++ {
		for j := len(A) - 1; j > i; j-- {
			n := Node{A[i], A[j], float64(A[i])/float64(A[j])}
			if h.Len() == K {
				if n.val >= (*h)[0].val {
					break
				}
			}
			heap.Push(h, n)
			if h.Len() > K {
				heap.Pop(h)
			}
		}
	}
	v := heap.Pop(h).(Node)
	return []int{v.x, v.y}
}

type NHeap []Node

func (h NHeap) Len() int           { return len(h) }
func (h NHeap) Less(i, j int) bool { return h[i].val > h[j].val }
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
