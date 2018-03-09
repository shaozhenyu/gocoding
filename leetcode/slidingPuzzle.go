package main

import "fmt"

func main() {
	board := [][]int{{3,1,0}, {4,2,5}}
	fmt.Println(slidingPuzzle(board))
}

func slidingPuzzle(board [][]int) int {
	a, b := 0, 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				a, b = i, j
			}
		}
	}
	smallest := goNext(board, 0, a, b, 20)
	if smallest == 20 {
		return -1
	}
	return smallest
}

func goNext(board [][]int, step int, i, j int, smallest int) int {
	if step > smallest {
		return smallest
	}
	if isOk(board) {
		return step
	}
	if i > 0 {
		tmp := make([][]int, 2)
		tmp[0] = make([]int, 3)
		copy(tmp[0], board[0])
		tmp[1] = make([]int, 3)
		copy(tmp[1], board[1])
		tmp[i][j], tmp[i-1][j] = tmp[i-1][j], tmp[i][j]
		smallest = goNext(tmp, step + 1, i-1, j, smallest)
	}

	if i < 1 {
		tmp := make([][]int, 2)
		tmp[0] = make([]int, 3)
		copy(tmp[0], board[0])
		tmp[1] = make([]int, 3)
		copy(tmp[1], board[1])
		tmp[i][j], tmp[i+1][j] = tmp[i+1][j], tmp[i][j]
		smallest = goNext(tmp, step + 1, i+1, j, smallest)
	}

	if j > 0 {
		tmp := make([][]int, 2)
		tmp[0] = make([]int, 3)
		copy(tmp[0], board[0])
		tmp[1] = make([]int, 3)
		copy(tmp[1], board[1])
		tmp[i][j], tmp[i][j-1] = tmp[i][j-1], tmp[i][j]
		smallest = goNext(tmp, step + 1, i, j-1, smallest)
	}

	if j < 2 {
		tmp := make([][]int, 2)
		tmp[0] = make([]int, 3)
		copy(tmp[0], board[0])
		tmp[1] = make([]int, 3)
		copy(tmp[1], board[1])
		tmp[i][j], tmp[i][j+1] = tmp[i][j+1], tmp[i][j]
		smallest = goNext(tmp, step + 1, i, j+1, smallest)
	}
	return smallest
}

func isOk(board [][]int) bool {
	val := 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != val%6 {
				return false
			} else {
				val++
			}
		}
	}
	return true
}
