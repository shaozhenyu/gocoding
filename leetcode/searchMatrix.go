package main

import "fmt"

func main() {
	matrix := [][]int{[]int{1}}
	fmt.Println(searchMatrix(matrix, 2))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	lo, li := 0, len(matrix) * len(matrix[0])
	for lo <= li {
		mid := lo + (li - lo)/2
		midi, midj := mid/len(matrix[0]), mid%len(matrix[0])
		if midi >= len(matrix) {
			return false
		}
		if matrix[midi][midj] < target {
			lo = mid + 1
		} else if matrix[midi][midj] > target {
			li = mid - 1
		} else {
			return true
		}
	}
	return false
}
