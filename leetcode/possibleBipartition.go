package main

import "fmt"

func main() {
	N := 10
	dislikes := [][]int{[]int{4, 7}, []int{4, 8}, []int{5, 6},[]int{1, 6},[]int{3,7},[]int{2,5},[]int{5,8},[]int{1,2},[]int{4,9},[]int{6,10},[]int{8,10},[]int{3,6},[]int{2,10},[]int{9,10},[]int{3,9},[]int{2,3},[]int{1,9},[]int{4,6},[]int{5,7}}
	fmt.Println(possibleBipartition(N, dislikes))
}

// [[4,7],[4,8],[5,6],[1,6],[3,7],[2,5],[5,8],[1,2],[4,9],[6,10],[8,10],[3,6],[2,10],[9,10],[3,9],[2,3],[1,9],[4,6],[5,7],[3,8],[1,8],[1,7],[2,4]]

func possibleBipartition(N int, dislikes [][]int) bool {
	m := make([][]int, N+1)
	visited := make([]bool, N+1)
	for i := 1; i <= N; i++ {
		m[i] = make([]int, 0, N)
	}
	for i := 0; i < len(dislikes); i++ {
		m[dislikes[i][0]] = append(m[dislikes[i][0]], dislikes[i][1])
		m[dislikes[i][1]] = append(m[dislikes[i][1]], dislikes[i][0])
	}
	a := make(map[int]bool)
	b := make(map[int]bool)

	var dfs func(idx int, a, b map[int]bool) bool
	dfs = func(idx int, a, b map[int]bool) bool {
		t := m[idx]
		for i := 0; i < len(t); i++ {
			if visited[t[i]] {
				continue
			}
			if a[t[i]] {
				return false
			}
			visited[t[i]] = true
			b[t[i]] = true
			if !dfs(t[i], b, a) {
				return false
			}
		}
		return true
	}

	for i := 1; i <= N; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		if a[i] || !b[i] {
			a[i] = true
			if !dfs(i, a, b) {
				return false
			}
		} else {
			b[i] = true
			if !dfs(i, b, a) {
				return false
			}
		}
	}
	return true
}
