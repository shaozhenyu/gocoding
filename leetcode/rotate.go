package main

import "fmt"

func main() {
	s := [][]int{[]int{1,2,3,4,5}, []int{6,7,8,9,10}, []int{11,12,13,14,15}, []int{16,17,18,19,20}, []int{21,22,23,24,25}}
	fmt.Println(rotate(s))
}

func rotate(matrix [][]int) [][]int {
	i, j := 0, 0
	lenM := len(matrix)
	for i <= lenM/2 {
		for i < lenM - 1 - j {
			tmp := matrix[i][j]
			i1 := j
			j1 := lenM - 1 - i
			for !(i1 == i && j1 == j) {
				val := tmp
				tmp = matrix[i1][j1]
				matrix[i1][j1] = val
				i1, j1 = j1, lenM - 1 - i1
				// fmt.Println(matrix, tmp, i, j)
			}
			val := tmp
			tmp = matrix[i1][j1]
			matrix[i1][j1] = val
			i++
		}
		j++
		i = j
		fmt.Println(i, j)
	}
	return matrix
}
