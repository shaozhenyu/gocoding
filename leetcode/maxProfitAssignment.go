package main

import (
	"fmt"
	"sort"
)

func main() {
	difficulty := []int{2,4,6,8,10}
	profit := []int{10,20,30,40,50}
	worker := []int{4,5,6,7}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))
}

type Node struct {
	d, p int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := make([]Node, len(profit))
	for i := 0; i < len(profit); i++ {
		n[i] = Node{difficulty[i], profit[i]}
	}
	sort.Slice(n, func(i, j int) bool {
		if n[i].p == n[j].p {
			return n[i].d < n[j].d
		}
		return n[i].p > n[j].p
	})
	ret := 0
	for i := 0; i < len(worker); i++ {
		for j := 0; j < len(n); j++ {
			if worker[i] >= n[j].d {
				ret += n[j].p
				break
			}
		}
	}
	return ret
}
