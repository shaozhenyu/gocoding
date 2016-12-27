package main

func main() {
	a := [][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}
	println(islandPerimeter(a))
}

// a left b up c right d down
func isConnect(grid [][]int, i, j, lenI, lenJ int) int {
	a, b, c, d := 0, 0, 0, 0
	left := j - 1
	up := i - 1
	right := j + 1
	down := i + 1
	if left >= 0 {
		a = grid[i][left]
	}
	if up >= 0 {
		b = grid[up][j]
	}
	if right <= lenJ {
		c = grid[i][right]
	}
	if down <= lenI {
		d = grid[down][j]
	}
	return (4 - a - b - c - d)
}

func islandPerimeter(grid [][]int) int {
	length := 0
	len1 := len(grid)
	for i := 0; i < len1; i++ {
		len2 := len(grid[i])
		for j := 0; j < len2; j++ {
			if grid[i][j] == 1 {
				length += isConnect(grid, i, j, len1-1, len2-1)
			}
		}
	}
	return length;
}
