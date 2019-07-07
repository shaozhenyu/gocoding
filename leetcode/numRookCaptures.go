package main

import "fmt"

func main() {

}

func numRookCaptures(board [][]byte) int {
	rx, ry := -1, -1
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				rx, ry = i, j
				break
			}
		}
		if rx != -1 {
			break
		}
	}
	count := 0
	for i := rx - 1; i >= 0; i-- {
		if board[i][ry] == 'B' {
			break
		} else if board[i][ry] == 'p' {
			count++
			break
		}
	}
	for i := rx + 1; i < len(board); i++ {
		if board[i][ry] == 'B' {
			break
		} else if board[i][ry] == 'p' {
			count++
			break
		}
	}
	for j := ry - 1; j >= 0; j-- {
		if board[rx][j] == 'B' {
			break
		} else if board[rx][j] == 'p' {
			count++
			break
		}
	}
	for j := ry + 1; j < len(board[rx]); j++ {
		if board[rx][j] == 'B' {
			break
		} else if board[rx][j] == 'p' {
			count++
			break
		}
	}
	return count
}
