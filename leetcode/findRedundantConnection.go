package main

import (
	"fmt"
)

func main() {
	edges := [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}
	edges = [][]int{[]int{1, 3}, []int{3, 4}, []int{1, 5}, []int{3, 5}, []int{2, 3}}
	fmt.Println(findRedundantConnection(edges))
}

func findRedundantConnection(edges [][]int) []int {
	union := make([]int, len(edges)+1)
	for i := 1; i < len(union); i++ {
		union[i] = i
	}

	ret := 0
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		if union[a] == union[b] {
			ret = i
		}
		idA := union[a]
		for i := 1; i < len(union); i++ {
			if union[i] == idA {
				union[i] = union[b]
			}
		}
	}
	return edges[ret]
}
