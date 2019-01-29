package main

import "fmt"

func main() {
	matrix := [][]int{[]int{1,0,1},[]int{0,-2,3}}
	k := 0
	matrix = [][]int{[]int{2, 2, -1}}
	fmt.Println(maxSumSubmatrix(matrix, k))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	smatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		smatrix[i] = make([]int, len(matrix[i]))
		sl := 0
		for j := 0; j < len(matrix[i]); j++ {
			sl += matrix[i][j]
			smatrix[i][j] += sl
			if i - 1 >= 0 {
				smatrix[i][j] += smatrix[i-1][j]
			}
			if smatrix[i][j] == k {
				return k
			}
		}
	}
	fmt.Println(smatrix)
	max := -1 << 31
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			for x := 0; x <= i; x++ {
				for y := 0; y <= j; y++ {
					sum := smatrix[i][j]
					if x - 1 >= 0 {
						sum -= smatrix[x-1][j]
					}
					if y - 1 >= 0 {
						sum -= smatrix[i][y-1]
					}
					if x - 1 >= 0 && y - 1 >= 0 {
						sum += smatrix[x-1][y-1]
					}
					if sum == k {
						return k
					}
					if sum < k && sum > max {
						max = sum
					}
				}
			}
		}
	}
	return max
}

