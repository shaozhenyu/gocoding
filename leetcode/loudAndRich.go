package main

import "fmt"

func main() {
	richer := [][]int{[]int{1,0},[]int{2,1},[]int{3,1},[]int{3,7},[]int{4,3},[]int{5,3},[]int{6,3}}
	quiet := []int{3,2,5,4,6,1,7,0}
	fmt.Println(loudAndRich(richer, quiet))
}

func loudAndRich(richer [][]int, quiet []int) []int {
	grid := make([][]int, len(quiet))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 0)
	}
	for i := 0; i < len(richer); i++ {
		grid[richer[i][1]] = append(grid[richer[i][1]], richer[i][0])
	}
	visited := make([]bool, len(quiet))
	ret := make([]int, len(quiet))
	m := make(map[int]int)
	for i := 0; i < len(quiet); i++ {
		dfs(i, grid, quiet, ret, visited)
		m[quiet[i]] = i
	}
	for i := 0; i < len(ret); i++ {
		ret[i] = m[ret[i]]
	}
	return ret
}

func dfs(idx int, grid [][]int, quiet, ret []int, visited []bool) int {
	if visited[idx] {
		return ret[idx]
	}
	min := quiet[idx]
	for i := 0; i < len(grid[idx]); i++ {
		q := dfs(grid[idx][i], grid, quiet, ret, visited)
		if q < min {
			min = q
		}
	}
	ret[idx] = min
	visited[idx] = true
	return min
}


