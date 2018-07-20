package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	queen := initQueens(n)
	return backtrack(0, queen, 0, n)
}

func backtrack(ret int, queen [][]byte, i, n int) int {
	if i == n {
		return ret + 1
	}
	for j := 0; j < n; j++ {
		if checkQueen(queen, i, j, n) {
			queen[i][j] = 'Q'
			ret = backtrack(ret, queen, i+1, n)
			queen[i][j] = '.'
		}
	}
	return ret
}

func getQueen(queen [][]byte) []string {
	ret := make([]string, len(queen))
	for i := 0; i < len(queen); i++ {
		ret[i] = string(queen[i])
	}
	return ret
}

func checkQueen(queen [][]byte, ii, jj, n int) bool {
	for i := 0; i < n; i++ {
		if queen[i][jj] == 'Q' {
			return false
		}
	}
	for j := 0; j < n; j++ {
		if queen[ii][j] == 'Q' {
			return false
		}
	}
	for i, j := ii, jj; i < n && j < n; i, j = i+1, j+1 {
		if queen[i][j] == 'Q' {
			return false
		}
	}
	for i, j := ii, jj; i >= 0 && j < n; i, j = i-1, j+1 {
		if queen[i][j] == 'Q' {
			return false
		}
	}
	for i, j := ii, jj; i < n && j >= 0; i, j = i+1, j-1 {
		if queen[i][j] == 'Q' {
			return false
		}
	}
	for i, j := ii, jj; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if queen[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func initQueens(n int) [][]byte {
	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		b[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			b[i][j] = '.'
		}
	}
	return b
}
