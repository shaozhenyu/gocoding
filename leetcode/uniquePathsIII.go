package main

import "fmt"

func main() {
	g := [][]int{[]int{1,0,0,0},[]int{0,0,0,0},[]int{0,0,2,-1}}
	fmt.Println(uniquePathsIII(g))
}

var step = [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}

func uniquePathsIII(grid [][]int) int {
	var startI, startJ, endI, endJ, need int
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				need += 1
			} else if grid[i][j] == 1 {
				startI, startJ = i, j
			} else if grid[i][j] == 2 {
				endI, endJ = i, j
			}
		}
	}
	need += 2
	total := 0
	visited[startI][startJ] = true
	backtrack(startI, startJ, endI, endJ, 1, need, grid, visited, &total)
	return total
}

func backtrack(x, y int, endI, endJ int, count, need int, grid [][]int, visited [][]bool, total *int) {
	if x == endI && y == endJ {
		if count == need {
			*total += 1
		}
		return
	}
	for _, s := range step {
		i := x + s[0]
		j := y + s[1]
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
			continue
		}
		if visited[i][j] || grid[i][j] == -1 {
			continue
		}
		visited[i][j] = true
		backtrack(i, j, endI, endJ, count + 1, need, grid, visited, total)
		visited[i][j] = false
	}
}
