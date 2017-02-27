package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := [][]byte{"..9748...","7........",".2.1.9...","..7...24.",".64.1.59.",".98...3..","...8.3.2.","........6","...2759.."}
	solveSudoku(b)
	fmt.Println(b)
}

func solveSudoku(board [][]byte)  {
	all := [9]int{9,9,9,9,9,9,9,9,9}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//fmt.Println(string(board[i][j]))
			num, err := strconv.Atoi(string(board[i][j]))
			if err == nil {
				all[num-1] -= 1
			} else {
				board[i][j] = byte(1)
			}
		}
	}
	fmt.Println(all)
}

