package main

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{[]int{1, 1}, []int{1, 2}}, 0, 0, 3))
}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	visited := make([][]bool, len(grid))
	ret := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
		ret[i] = make([]int, len(grid[i]))
	}
	dfs(visited, ret, grid, r0, c0, grid[r0][c0], color)
	fmt.Println(ret)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if ret[i][j] == 0 {
				ret[i][j] = grid[i][j]
			}
		}
	}
	return ret
}

func dfs(visited [][]bool, ret, grid [][]int, x, y, old, color int) {
	if visited[x][y] {
		return
	}
	visited[x][y] = true
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
		ret[x][y] = color
	}
	if x - 1 >= 0 && grid[x-1][y] == old {
		dfs(visited, ret, grid, x-1, y, old, color)
	} else {
		ret[x][y] = color
	}

	if x + 1 < len(grid) &&  grid[x+1][y] == old {
		dfs(visited, ret, grid, x+1, y, old, color)
	} else {
		ret[x][y] = color
	}
	if y - 1 >= 0 && grid[x][y-1] == old {
		dfs(visited, ret, grid, x, y-1, old, color)
	} else {
		ret[x][y] = color
	}
	if y + 1 < len(grid[x]) && grid[x][y+1] == old {
		dfs(visited, ret, grid, x, y+1, old, color)
	} else {
		ret[x][y] = color
	}
}
