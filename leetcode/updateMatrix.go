package main

import "fmt"

func main() {
	matrix := [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}
	fmt.Println(updateMatrix(matrix))
}

var step = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, 1}, []int{0, -1}}

func updateMatrix(matrix [][]int) [][]int {
	visited := make([][]bool, len(matrix))
	ret := make([][]int, len(matrix))
	queue := make([][]int, len(matrix) * len(matrix[0]))
	idx := 0

	for i := 0; i < len(matrix); i++ {
		ret[i] = make([]int, len(matrix[i]))
		visited[i] = make([]bool, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				ret[i][j] = 0
				visited[i][j] = true
				queue[idx] = []int{i, j}
				idx++
			}
		}
	}
	start := 0
	for start < idx {
		now := queue[start]
		start++
		for _, s := range step {
			x, y := now[0] + s[0], now[1] + s[1]
			if x < 0 || x >= len(ret) || y < 0 || y >= len(ret[0]) {
				continue
			}
			if visited[x][y] {
				continue
			}
			visited[x][y] = true
			ret[x][y] = 1 + ret[now[0]][now[1]]
			queue[idx] = []int{x, y}
			idx++
		}
	}
	return ret
}
