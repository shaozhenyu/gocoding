package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	k := 2
	W := 0
	Profits := []int{1, 2, 3}
	Capital := []int{0, 1, 1}
	fmt.Println(findMaximizedCapital(k, W, Profits, Capital))
}

type Node struct {
	profit  int
	capital int
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	n := make([]Node, len(Profits))
	for i := 0; i < len(Profits); i++ {
		n[i] = Node{Profits[i], Capital[i]}
	}
	sort.Slice(n, func(i, j int) bool {
		if n[i].profit == n[j].profit {
			return n[i].capital < n[j].capital
		}
		return n[i].profit > n[j].profit
	})
	l := list.New()
	for i := 0; i < len(n); i++ {
		l.PushBack(n[i])
	}
	for k > 0 {
		for e := l.Front(); e != nil; e = e.Next() {
			node := e.Value.(Node)
			if node.capital <= W {
				W += node.profit
				l.Remove(e)
				break
			}
		}
		k--
	}
	return W
}
