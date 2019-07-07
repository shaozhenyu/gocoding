package main

import (
	"fmt"
	"sort"
)

func main() {
	workers := [][]int{[]int{0, 0}, []int{2, 1}}
	bikes := [][]int{[]int{1,2},[]int{3,3}}
	fmt.Println(assignBikes(workers, bikes))
}

type Node struct {
	Manhattan int
	workIdx int
	bikeIdx int
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	node := make([]Node, len(workers) * len(bikes))
	idx := 0
	for i := 0; i < len(workers); i++ {
		for j := 0; j < len(bikes); j++ {
			node[idx] = Node{mh(workers[i], bikes[j]), i, j}
			idx++
		}
	}
	sort.Slice(node, func(i, j int) bool {
		if node[i].Manhattan == node[j].Manhattan {
			if node[i].workIdx == node[j].workIdx {
				return node[i].bikeIdx < node[j].bikeIdx
			}
			return node[i].workIdx < node[j].workIdx
		}
		return node[i].Manhattan < node[j].Manhattan
	})
	fmt.Println(node)
	work := make(map[int]bool)
	bike := make(map[int]bool)
	ret := make([]int, len(workers))
	for i := 0; i < len(node); i++ {
		if work[node[i].workIdx] || bike[node[i].bikeIdx] {
			continue
		}
		ret[node[i].bikeIdx] = node[i].workIdx
		work[node[i].workIdx] = true
		bike[node[i].bikeIdx] = true
	}
	return ret
}

func mh(x, y []int) int {
	return abs(x[0] - y[0]) + abs(x[1] - y[1])
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
