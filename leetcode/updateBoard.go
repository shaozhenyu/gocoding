package main

import "fmt"

func main() {
	board := [][]byte{{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}}
	click := []int{3, 0}
	fmt.Println(updateBoard(board, click))
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	x, y := click[0], click[1]

	count := 0
	if x-1 >= 0 {
		count += isMOrX(x-1, y, board)
	}
	if x+1 < len(board) {
		count += isMOrX(x+1, y, board)
	}
	if y-1 >= 0 {
		count += isMOrX(x, y-1, board)
	}
	if y+1 < len(board[0]) {
		count += isMOrX(x, y+1, board)
	}
	if x-1 >= 0 && y-1 >= 0 {
		count += isMOrX(x-1, y-1, board)
	}
	if x-1 >= 0 && y+1 < len(board[0]) {
		count += isMOrX(x-1, y+1, board)
	}
	if x+1 < len(board) && y-1 >= 0 {
		count += isMOrX(x+1, y-1, board)
	}
	if x+1 < len(board) && y+1 < len(board[0]) {
		count += isMOrX(x+1, y+1, board)
	}

	if count == 0 {
		board[x][y] = 'B'
		if x-1 >= 0 && board[x-1][y] == 'E' {
			updateBoard(board, []int{x - 1, y})
		}
		if x+1 < len(board) && board[x+1][y] == 'E' {
			updateBoard(board, []int{x + 1, y})
		}
		if y-1 >= 0 && board[x][y-1] == 'E' {
			updateBoard(board, []int{x, y - 1})
		}
		if y+1 < len(board[0]) && board[x][y+1] == 'E' {
			updateBoard(board, []int{x, y + 1})
		}
		if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1] == 'E' {
			updateBoard(board, []int{x - 1, y - 1})
		}
		if x-1 >= 0 && y+1 < len(board[0]) && board[x-1][y+1] == 'E' {
			updateBoard(board, []int{x - 1, y + 1})
		}
		if x+1 < len(board) && y-1 >= 0 && board[x+1][y-1] == 'E' {
			updateBoard(board, []int{x + 1, y - 1})
		}
		if x+1 < len(board) && y+1 < len(board[0]) && board[x+1][y+1] == 'E' {
			updateBoard(board, []int{x + 1, y + 1})
		}
	} else {
		board[x][y] = byte(count)
	}

	return board
}

func isMOrX(x, y int, board [][]byte) int {
	if board[x][y] == 'M' || board[x][y] == 'X' {
		return 1
	}

	return 0
}
