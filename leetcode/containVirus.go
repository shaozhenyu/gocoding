package main

import "fmt"

func main() {
	grid := [][]int{[]int{1,1,1}, []int{1,0,1}, []int{1,1,1}}
	fmt.Println(containVirus(grid))
}

func containVirus(grid [][]int) int {
	// 设置隔离区
	g := make([][]int, len(grid) * 2 - 1)
	for i := 0; i < len(g); i++ {
		g[i] = make([]int, len(grid[0]) * 2 - 1)
		for j := 0; j < len(g[i]); j++ {
			if i%2 == 0 && j%2 == 0 {
				g[i][j] = grid[i/2][j/2]
			} else {
				g[i][j] = -1
			}
		}
	}
	fmt.Println(g)
	// 找到对安全区威胁最大的病毒群
	return 0
}
