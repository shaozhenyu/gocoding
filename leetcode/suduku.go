package main

import "fmt"

func main() {
	//broad := [][]byte{byte(1), byte(2), byte(3), byte(4)}
	broad := [][]byte{{46, 49, 50, 51, 52, 53, 53, 54, 55},
	{46, 49, 50, 52, 52, 53, 53, 54, 55},
	{46, 49, 50, 56, 52, 53, 53, 54, 55},
	{46, 49, 50, 50, 52, 53, 53, 54, 55},
	{46, 49, 50, 51, 52, 53, 53, 54, 55},
	{46, 49, 50, 57, 52, 53, 53, 54, 55},
	{46, 49, 50, 54, 52, 53, 53, 54, 55},
	{46, 49, 50, 51, 52, 53, 53, 54, 55},
	{46, 49, 50, 55, 52, 53, 53, 54, 55}}

	if isValidSudoku(broad) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isValidOne(a [9]int) bool {
	for i := 0; i < 8; i++ {
		if a[i] == 10 {
			continue
		}
		for j := i+1; j < 9; j++ {
			if a[j] == 10 {
				continue
			} else if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {

	arr := [9][9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) == "." {
				arr[i][j] = 10
			} else if board[i][j] > 48 && board[i][j] < 58 {
				arr[i][j] = int(board[i][j]) -48
			} else {
				return false
			}
		}
	}

	b := [9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b[j] = arr[i][j]
		}
		if !isValidOne(b) {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b[j] = arr[j][i]
		}
		if !isValidOne(b) {
			return false
		}
	}

	for n := 0; n < 3; n++ {
		for m := 0; m < 3; m++ {
			ii := 3 * n
			jj := 3 * m
			k := 0
			for i := ii; i < 3*(n+1); i++ {
				for j := jj; j < 3*(m+1); j++ {
					b[k] = arr[i][j]
					k++
				}
			}

			if !isValidOne(b) {
				return false
			}
		}
	}

	return true
}
