package main

import "fmt"

func main() {
	board := [][]int{[]int{-1,-1,-1,-1,-1,-1},[]int{-1,-1,-1,-1,-1,-1},[]int{-1,-1,-1,-1,-1,-1},[]int{-1,35,-1,-1,13,-1},[]int{-1,-1,-1,-1,-1,-1},[]int{-1,15,-1,-1,-1,-1}}
	board = [][]int{[]int{-1,4,-1},[]int{6,2,6},[]int{-1,3,-1}}
	board = [][]int{[]int{-1,-1,19,10,-1},[]int{2,-1,-1,6,-1},[]int{-1,17,-1,19,-1},[]int{25,-1,20,-1,-1},[]int{-1,-1,-1,-1,15}}
	fmt.Println(snakesAndLadders(board))
}

func snakesAndLadders(board [][]int) int {
	minReach := make([][]int, len(board))
	for i := 0; i < len(minReach); i++ {
		minReach[i] = make([]int, len(board[0]))
	}
	minReach[len(board)-1][0] = 1
	min := 2 << 31
	dfs(1, 0, len(board), &min, board, minReach)
	if min == 2 << 31 {
		return -1
	}
	return min
}

func dfs(idx, step, n int, min *int, board [][]int, minReach [][]int) {
	if step >= *min {
		return
	}
	if idx >= n * n {
		if *min > step {
			*min = step
		}
		return
	}
	x, y := getIdx(idx, n)
	if board[x][y] != -1 {
		idx = board[x][y]
	}
	if idx >= n * n {
		if *min > step {
			*min = step
		}
		return
	}
	x, y = getIdx(idx, n)
	if minReach[x][y] > 0 && minReach[x][y] <= step {
		return
	}
	minReach[x][y] = step
	for i := 1; i <= 6; i++ {
		dfs(idx + i, step + 1, n, min, board, minReach)
	}
	return
}

func getIdx(idx int, n int) (i int,j int) {
	i = n - 1 - (idx-1)/n
	if (n - 1 - i)%2 == 0 {
		j = (idx%n + n - 1)%n
	} else {
		j = (n - idx%n)%n
	}
	return
}
