package main

import "fmt"

func main() {
	matrix := [][]byte{
		[]byte{'1', '1', '1', '1', '1', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1', '1', '1', '0'},
		[]byte{'1', '1', '1', '1', '1', '1', '1', '0'},
		[]byte{'1', '1', '1', '1', '1', '0', '0', '0'},
		[]byte{'0', '1', '1', '1', '1', '0', '0', '0'}}
	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	fmt.Println(matrix)
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			endI := i + 1
			for endI < len(matrix) {
				if matrix[endI][j] == '0' {
					break
				}
				endI++
			}
			endJ := j + 1
			for endJ < len(matrix[i]) {
				if matrix[i][endJ] == '0' {
					break
				}
				endJ++
			}
			count := checkRectangle(matrix, i, j, endI, endJ)
			fmt.Println(count, i, j)
			if count > max {
				max = count
			}
		}
	}
	return max
}

func checkRectangle(matrix [][]byte, startI, startJ, endI, endJ int) int {
	max := 0
	for i := startI; i < endI; i++ {
		for j := startJ; j < endJ; j++ {
			if matrix[i][j] == '0' {
				endJ = j
				break
			}
		}
		if (i - startI + 1) * (endJ - startJ) > max {
			max = (i - startI + 1) * (endJ - startJ)
		}
	}
	return max
}
