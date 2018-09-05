package main

import "fmt"

func main() {
	val := [][]byte{[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}, []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}, []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}, []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}, []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}, []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(val)
}


func solveSudoku(board [][]byte) {
	a := make([][]bool, 9)
	b := make([][]bool, 9)
	c := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		a[i] = make([]bool, 9)
		b[i] = make([]bool, 9)
		c[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				a[i][board[i][j]-'1'] = true
				b[j][board[i][j]-'1'] = true
				c[i/3*3+j/3][board[i][j]-'1'] = true
			}
		}
	}
	backtrace(board, a, b, c, 0)
	fmt.Println("after: ", board)
}

var (
	all = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
)
func backtrace(board [][]byte, a, b, c [][]bool, index int) bool {
	if index == 81 {
		return true
	}
	var start int
	for start = index; start < 81; start++ {
		i, j := start/9, start%9
		if board[i][j] == '.' {
			break
		}
	}
	if start >= 81 {
		return true
	}
	i, j := start/9, start%9
	for x := 0; x < 9; x++ {
		v := all[x] - '1'
		if !a[i][v] && !b[j][v] && !c[i/3*3+j/3][v] {
			board[i][j] = all[x]
			a[i][v] = true
			b[j][v] = true
			c[i/3*3+j/3][v] = true
			p := i
			q := j + 1
			if q == 9 {
				q = 0
				p++
			}
			ok := backtrace(board, a, b, c, start+1)
			if ok {
				return true
			}
			board[i][j] = '.'
			a[i][v] = false
			b[j][v] = false
			c[i/3*3+j/3][v] = false
		}
	}
	return false
}
