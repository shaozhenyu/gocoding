package main

import "fmt"

func main() {
	fmt.Println(leadsToDestination(4, [][]int{[]int{0, 1}, []int{0, 2}, []int{1,2}, []int{1, 3},[]int{2,3}}, 0, 3))
}

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	es := make([][]int, n)
	for i := 0; i < n; i++ {
		es[i] = make([]int, 0)
	}
	for i := 0; i < len(edges); i++ {
		if edges[i][0] == edges[i][1] {
			return false
		}
		if edges[i][0] == destination {
			return false
		}
		es[edges[i][0]] = append(es[edges[i][0]], edges[i][1])
	}
	total := make(map[int]bool)
	for i := 0 ; i < len(es[source]); i++ {
		m := make(map[int]bool)
		m[source] = true
		if !dfs(es[source][i],destination, es, 0, m, total) {
			return false
		}
	}
	return true
}

func dfs(now, destination int, es [][]int, count int, m map[int]bool, total map[int]bool) bool {
	if now != destination && len(es[now]) == 0 {
		return false
	}
	k := 0
	for i := 0; i < len(es[now]); i++ {
		if es[now][i] == destination {
			k++
			continue
		}
		if m[es[now][i]] {
			return false
		}
		m[es[now][i]] = true
		if !dfs(es[now][i], destination, es, 0, m, total) {
			return false
		} else {
			k++
		}
		m[es[now][i]] = false
	}
	return k == len(es[now])
}
