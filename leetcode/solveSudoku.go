package main

import "fmt"

func main() {
	val := [][]byte{[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}, []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}, []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}, []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}, []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}, []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(val)
}

var (
	all = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
)

func solveSudoku(board [][]byte) {
	a := make([]map[byte]bool, 9)
	b := make([]map[byte]bool, 9)
	c := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		a[i] = make(map[byte]bool)
		b[i] = make(map[byte]bool)
		c[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				a[i][board[i][j]] = true
				b[j][board[i][j]] = true
				c[i/3*3+j/3][board[i][j]] = true
			}
		}
	}
	backtrace(board, a, b, c, 0, 0)
}

func backtrace(board [][]byte, a, b, c []map[byte]bool, m, n int) bool {
	for i := m; i < 9; i++ {
		for j := n; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			for x := 0; x < 9; x++ {
				v := all[x]
				if !a[i][v] && !b[j][v] && !c[i/3*3+j/3][v] {
					board[i][j] = v
					a[i][v] = true
					b[j][v] = true
					c[i/3*3+j/3][v] = true
					q := j + 1
					p := i
					if q == 9 {
						q = 0
						p++
					}
					ok := backtrace(board, a, b, c, p, q)
					if ok {
						return ok
					}
					board[i][j] = '.'
					a[i][v] = false
					b[j][v] = false
					c[i/3*3+j/3][v] = false
				}
			}
		}
	}
	return false
}
