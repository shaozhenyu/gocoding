package main

import "fmt"

func main() {
	grid := []string{"\\/\\ "," /\\/"," \\/ ","/ / "}
	fmt.Println(regionsBySlashes(grid))
}

func regionsBySlashes(grid []string) int {
	fmt.Println(grid)
	if len(grid) == 0 {
		return 0
	}
	N := len(grid) + 1
	u := make([]int, N*N)
	count := 0
	for i := 0; i < N*N; i++ {
		u[i] = i
	}
	for j := 0; j < N-1; j++ {
		union(j, j+1, u)
	}
	for i := 0; i < N-1; i++ {
		union(i*N, (i+1)*N, u)
	}
	for j := 0; j < N-1; j++ {
		union(N*(N-1)+j, N*(N-1)+j+1, u)
	}
	for i := 0; i < N-1; i++ {
		union(N-1+i*N, N-1+(i+1)*N, u)
	}
	fmt.Println(u)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '/' {
				if union(i*N+j+1, (i+1)*N+j, u) {
					count++
				}
			} else if grid[i][j] == '\\' {
				if union(i*N+j, (i+1)*N+j+1, u) {
					count++
				}
			}
			fmt.Println(i, j, u)
		}
	}
	return count + 1
}

func find(x int, u []int) int {
	if x != u[x] {
		u[x] = find(u[x], u)
	}
	return u[x]
}

func union(x, y int, u []int) bool {
	fx, fy := find(x, u), find(y, u)
	if fx != fy {
		// u[y] = fx
		for i := 0; i < len(u); i++ {
			if u[i] == fy {
				u[i] = fx
			}
		}
		return false
	}
	return true
}
