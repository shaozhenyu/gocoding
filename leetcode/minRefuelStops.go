package main

import (
	"fmt"
	"container/heap"
)

func main() {
	target := 100
	startFuel := 50
	stations := [][]int{[]int{25,25},[]int{50,50}}
	fmt.Println(minRefuelStops(target, startFuel, stations))
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	fuel := startFuel
	if fuel >= target {
		return 0
	}
	ret := 0
	total := &IntHeap{}
	heap.Init(total)
	for i := 0; i < len(stations); i++ {
		for fuel < stations[i][0] {
			if total.Len() == 0 {
				return -1
			}
			fuel += heap.Pop(total).(int)
			ret++
		}
		heap.Push(total, stations[i][1])
	}
	for fuel < target {
		if total.Len() == 0 {
			return -1
		}
		fuel += heap.Pop(total).(int)
		ret++
	}
	return ret
}

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
