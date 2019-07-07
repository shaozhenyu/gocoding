package main

import (
	"fmt"
	"sort"
)

func main() {

}

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return abs(costs[i][0] - costs[i][1]) >abs(costs[j][0] - costs[j][1])
	})
	a, b, n := 0, 0, len(costs)/2
	cost := 0
	for i := 0; i < len(costs); i++ {
		if costs[i][0] <= costs[i][1] {
			if a < n {
				cost += costs[i][0]
				a++
			} else {
				cost += costs[i][1]
				b++
			}
		} else {
			if b < n {
				cost += costs[i][1]
				b++
			} else {
				cost += costs[i][0]
				a++
			}
		}
	}
	return cost
}


func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
