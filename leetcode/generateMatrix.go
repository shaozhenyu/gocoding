package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	nn := n * n
	flag := 0
	var i, j int
	for num := 1; num <= nn; num++ {
		ret[i][j] = num
		switch flag {
		case 0:
			if j < n-1 && ret[i][j+1] == 0 {
				j++
			} else {
				flag = (flag + 1) % 4
				i++
			}
		case 1:
			if i < n-1 && ret[i+1][j] == 0 {
				i++
			} else {
				flag = (flag + 1) % 4
				j--
			}
		case 2:
			if j > 0 && ret[i][j-1] == 0 {
				j--
			} else {
				flag = (flag + 1) % 4
				i--
			}
		case 3:
			if i > 0 && ret[i-1][j] == 0 {
				i--
			} else {
				flag = (flag + 1) % 4
				j++
			}
		}

	}
	return ret
}
