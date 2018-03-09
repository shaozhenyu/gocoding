package main

import "fmt"

var m map[int]int

func main() {
	m = make(map[int]int)
	initMap([][]int{{1,2,3},{4,5,0}}, 0, 1, 2)
	board := [][]int{{3,2,4}, {1,5,0}}
	fmt.Println(slidingPuzzle(board))
}

func initMap(board [][]int, step int, i, j int) {
	val := parseInt(board)
	if old, ok := m[val]; ok {
		if step >= old {
			return
		}
	}
	m[val] = step
	if i > 0 {
		tmp := make([][]int, 2)
		tmp[0] = make([]int, 3)
		copy(tmp[0], board[0])
		tmp[1] = make([]int, 3)
		copy(tmp[1], board[1])
		tmp[i][j], tmp[i-1][j] = tmp[i-1][j], tmp[i][j]
		initMap(tmp, step + 1, i-1, j)
	}

	if i < 1 {
		tmp := make([][]int, 2)
		tmp[0] = make([]int, 3)
		copy(tmp[0], board[0])
		tmp[1] = make([]int, 3)
		copy(tmp[1], board[1])
		tmp[i][j], tmp[i+1][j] = tmp[i+1][j], tmp[i][j]
		initMap(tmp, step + 1, i+1, j)
	}

	if j > 0 {
		tmp := make([][]int, 2)
		tmp[0] = make([]int, 3)
		copy(tmp[0], board[0])
		tmp[1] = make([]int, 3)
		copy(tmp[1], board[1])
		tmp[i][j], tmp[i][j-1] = tmp[i][j-1], tmp[i][j]
		initMap(tmp, step + 1, i, j-1)
	}

	if j < 2 {
		tmp := make([][]int, 2)
		tmp[0] = make([]int, 3)
		copy(tmp[0], board[0])
		tmp[1] = make([]int, 3)
		copy(tmp[1], board[1])
		tmp[i][j], tmp[i][j+1] = tmp[i][j+1], tmp[i][j]
		initMap(tmp, step + 1, i, j+1)
	}

}

func parseInt(board [][]int) int {
	ret := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			ret = board[i][j] + ret * 10
		}
	}
	return ret
}

func slidingPuzzle(board [][]int) int {
	val := parseInt(board)
	if step, ok := m[val]; ok {
		return step
	}
	return -1
}
