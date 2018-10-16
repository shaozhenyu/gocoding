package main

import "fmt"

func main() {
	// v := [][]int{[]int{1,2,3,4}, []int{5,6,7,8}, []int{9,10,11,12}}
	v := [][]int{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15}, []int{16, 17, 18, 19, 20}, []int{21, 22, 23, 24, 25}}
	fmt.Println(spiralOrder(v))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	fmt.Println(matrix)
	spiral := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	times := []int{len(matrix[0]) - 1, len(matrix) - 1, len(matrix[0]) - 1, len(matrix) - 2}
	i, j := 0, 0
	index := 0
	size := len(matrix) * len(matrix[0])
	ret := make([]int, size)
	ret[index] = matrix[i][j]
	index++
	for index < size {
		for x := 0; x < 4; x++ {
			if index >= size {
				break
			}
			t := 1
			for index < size && t <= times[x] {
				i += spiral[x][0]
				j += spiral[x][1]
				fmt.Println(i, j)
				ret[index] = matrix[i][j]
				t++
				index++
			}
			times[x] -= 2
			fmt.Println("ret: ", ret, index)
		}
		times[0] = times[2] + 1
		fmt.Println("times : ", times)
	}
	return ret
}
