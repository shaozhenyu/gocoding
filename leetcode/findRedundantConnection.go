package main

import (
	"fmt"
)

func quickFind(edges [][]int) []int {
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

func quickUnion(edges [][]int) []int {
	union := make([]int, len(edges)+1)
	for i := 1; i < len(union); i++ {
		union[i] = i
	}

	ret := 0
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]

		for a != union[a] {
			a = union[a]
		}

		for b != union[b] {
			b = union[b]
		}

		if a == b {
			ret = i
			continue
		}

		union[a] = b
	}
	return edges[ret]
}

func WeightedquickUnion(edges [][]int) []int {
	union := make([]int, len(edges)+1)
	size := make([]int, len(edges)+1)
	for i := 1; i < len(union); i++ {
		union[i] = i
		size[i] = 1
	}

	ret := 0
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]

		for a != union[a] {
			a = union[a]
		}
		union[edges[i][0]] = a

		for b != union[b] {
			b = union[b]
		}
		union[edges[i][1]] = b

		if a == b {
			ret = i
			continue
		}
		if size[a] > size[b] {
			size[a] += size[b]
			union[b] = a
		} else {
			size[b] += size[a]
			union[a] = b
		}

	}
	return edges[ret]
}

func main() {
	edges := [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}
	edges = [][]int{[]int{1, 3}, []int{3, 4}, []int{1, 5}, []int{3, 5}, []int{2, 3}}
	edges = [][]int{[]int{3, 4}, []int{1, 2}, []int{2, 4}, []int{3, 5}, []int{2, 5}}
	fmt.Println(quickUnion(edges))
}
