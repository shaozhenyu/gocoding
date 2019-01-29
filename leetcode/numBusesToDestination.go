package main

import "fmt"

func main() {
	routes := [][]int{[]int{1, 2, 7}, []int{3, 6, 7}}
	S := 1
	T := 6
	fmt.Println(numBusesToDestination(routes, S, T))
}
type Node struct {
	station int
	count int
}

func numBusesToDestination(routes [][]int, S int, T int) int {
	m := make(map[int]map[int]bool)
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			if _, ok := m[routes[i][j]]; !ok {
				m[routes[i][j]] = make(map[int]bool)
			}
			m[routes[i][j]][i] = true
		}
	}
	visitedBus := make(map[int]bool)
	visitedRoute := make(map[int]bool)
	q := make([]Node, 100000)
	start, end := 0, 0
	q[start] = Node{S, 0}
	for start <= end {
		now := q[start]
		if T == now.station {
			return now.count
		}
		if route, ok := m[now.station]; ok {
			for rs := range route {
				if visitedRoute[rs] {
					continue
				}
				visitedRoute[rs] = true
				for i := 0; i < len(routes[rs]); i++ {
					if visitedBus[routes[rs][i]] {
						continue
					}
					visitedBus[routes[rs][i]] = true
					end++
					q[end] = Node{routes[rs][i], now.count+1}
				}
			}
		}
		start++
	}
	return -1
}

