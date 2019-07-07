package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(allCellsDistOrder(1, 100, 0, 35))
}

type Node struct {
	val []int
	dist int
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	node := make([]Node, R*C)
	idx := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			node[idx] = Node{[]int{i, j}, dist(i, j, r0, c0)}
			idx++
		}
	}
	sort.Slice(node, func(i, j int) bool {
		return node[i].dist < node[j].dist
	})
	ret := make([][]int, len(node))
	for i := 0; i < len(node); i++ {
		ret[i] = node[i].val
	}
	return ret
}

func dist(r, c, r0, c0 int) int {
	return abs(r - r0) + abs(c - c0)
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
