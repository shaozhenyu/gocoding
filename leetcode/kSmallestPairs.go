package main

import (
	"fmt"
	"container/heap"
)

func main() {
	nums1 := []int{1,7,11}
	nums2 := []int{2,4,6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

// o(kn)
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) * len(nums2) < k {
		k = len(nums1) * len(nums2)
	}
	t := make([]int, len(nums1))
	ret := make([][]int, 0)
	for k > 0 {
		min := 2 << 31
		minIdx := 0
		for i := 0; i < len(nums1); i++ {
			if t[i] >= len(nums2) {
				continue
			}
			if nums1[i] + nums2[t[i]] < min {
				min = nums1[i] + nums2[t[i]]
				minIdx = i
			}
		}
		ret = append(ret, []int{nums1[minIdx], nums2[t[minIdx]]})
		t[minIdx]++
		k--
	}
	return ret
}

type Node struct {
	sum int
	i, j int
}

type NHeap []Node
func (h NHeap) Len() int           { return len(h) }
func (h NHeap) Less(i, j int) bool { return h[i].sum > h[j].sum }
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

func kSmallestPairs1(nums1 []int, nums2 []int, k int) [][]int {
	h := &NHeap{}
	heap.Init(h)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(h, Node{nums1[i] + nums2[j], nums1[i], nums2[j]})
			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}
	ret := make([][]int, k)
	for k > 0 {
		v := heap.Pop(h).(Node)
		ret[k-1] = []int{v.i, v.j}
		k--
	}
	return ret
}
