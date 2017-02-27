package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	ret := [][]string{}
	s := make([][]byte, n)
	for i := 0; i < n; i++ {
		str := ""
		for i := 0; i < n; i++ {
			str += "."
		}
		s[i] = []byte(str)
	}
	ret = backtrace(n, 0, s, ret)
	return ret
}

func backtrace(n int, row int, s [][]byte, ret [][]string) [][]string {
	if row == n {
		success := make([]string, n)
		for i := 0; i < n; i++ {
			success[i] = string(s[i])
		}
		ret = append(ret, success)
		return ret
	}
	for i := 0; i < n; i++ {
		if isValid(n, s, row, i) {
			s[row][i] = byte('Q')
			ret = backtrace(n, row+1, s, ret)
			s[row][i] = byte('.')
		}
	}
	return ret
}

func isValid(n int, s [][]byte, row int, column int) bool {
	Q := byte('Q')
	for i := 0; i < n; i++ {
		if s[i][column] == Q {
			return false
		}
	}
	for i, j := row+1, column-1; i < n && j >= 0; i, j = i+1, j-1 {
		if s[i][j] == Q {
			return false
		}
	}
	for i, j := row-1, column+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if s[i][j] == Q {
			return false
		}
	}
	for i, j := row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if s[i][j] == Q {
			return false
		}
	}
	for i, j := row+1, column+1; i < n && j < n; i, j = i+1, j+1 {
		if s[i][j] == Q {
			return false
		}
	}
	return true
}
