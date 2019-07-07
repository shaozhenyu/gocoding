package main

import "fmt"

func main() {
	points := [][]int{[]int{1,1},[]int{2,3},[]int{3,2}}
	points = [][]int{[]int{0,1},[]int{1,1},[]int{2,1}}
	fmt.Println(isBoomerang(points))
}

func isBoomerang(points [][]int) bool {
	if points[0][0] == points[1][0] && points[0][1] == points[1][1] {
		return false
	}
	if points[0][0] == points[2][0] && points[0][1] == points[2][1] {
		return false
	}
	if points[2][0] == points[1][0] && points[2][1] == points[1][1] {
		return false
	}
	if (points[0][0] - points[2][0]) * (points[1][1] - points[2][1]) == (points[1][0] - points[2][0]) * (points[0][1] - points[2][1]) {
		return false
	}
	return true
}

