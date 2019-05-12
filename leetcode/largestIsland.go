package main

import "fmt"

func main() {
	grid := [][]int{[]int{0, 0}, []int{0, 1}}
	fmt.Println(largestIsland(grid))
}

type Node struct {
	startX, startY int
	count int
}

var step = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}

func largestIsland(grid [][]int) int {
	n := make([][]Node, len(grid))
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		n[i] = make([]Node, len(grid[i]))
		visited[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if visited[i][j] {
				continue
			}
			if grid[i][j] == 1 {
				n[i][j] = Node{i, j, 0}
				tr(i, j, i, j, grid, n, visited)
			}
		}
	}
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				max = Max(max, n[n[i][j].startX][n[i][j].startY].count)
			} else {
				m := make(map[string]bool)
				count := 1
				for _, s := range step {
					x := i + s[0]
					y := j + s[1]
					if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
						continue
					}
					if grid[x][y] == 0 {
						continue
					}
					startX, startY := n[x][y].startX, n[x][y].startY
					if !m[fmt.Sprintf("%d-%d", startX, startY)] {
						count += n[startX][startY].count
						m[fmt.Sprintf("%d-%d", startX, startY)] = true
					}
				}
				fmt.Println(m, count)
				max = Max(max, count)
			}
		}
	}
	fmt.Println(n)
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func tr(startX, startY int, x, y int, grid [][]int, n [][]Node, visited [][]bool) {
	n[startX][startY].count++
	visited[x][y] = true
	for _, s := range step {
		i := x + s[0]
		j := y + s[1]
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			continue
		}
		if visited[i][j] || grid[i][j] == 0 {
			continue
		}
		n[i][j] = Node{startX, startY, 0}
		tr(startX, startY, i, j, grid, n, visited)
	}
}
