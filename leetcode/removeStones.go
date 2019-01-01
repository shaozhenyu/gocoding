package main

import "fmt"

func main() {
	stone := [][]int{[]int{0, 0}, []int{0, 1}, []int{1, 0}, []int{1, 2}, []int{2, 1}, []int{2, 2}}
	fmt.Println(removeStones(stone))
}

func removeStones(stones [][]int) int {
	visited := make(map[int]map[int]bool)
	xm := make(map[int][]int)
	ym := make(map[int][]int)
	for _, s := range stones {
		x, y := s[0], s[1]
		if _, ok := xm[x]; !ok {
			xm[x] = make([]int, 0)
		}
		xm[x] = append(xm[x], y)
		if _, ok := ym[y]; !ok {
			ym[y] = make([]int, 0)
		}
		ym[y] = append(ym[y], x)
		if _, ok := visited[x]; !ok {
			visited[x] = make(map[int]bool)
		}
		visited[x][y] = false
	}
	count := len(stones)
	for _, s := range stones {
		if  !visited[s[0]][s[1]] {
			dfs(s[0], s[1], xm, ym, visited)
			count--
		}
	}
	return count
}

func dfs(x, y int, xm, ym map[int][]int, visited map[int]map[int]bool) {
	visited[x][y] = true
	for i := 0; i < len(xm[x]); i++ {
		if !visited[x][xm[x][i]] {
			dfs(x, xm[x][i], xm, ym, visited)
		}
	}
	for j := 0; j < len(ym[y]); j++ {
		if !visited[ym[y][j]][y] {
			dfs(ym[y][j], y, xm, ym, visited)
		}
	}
}
