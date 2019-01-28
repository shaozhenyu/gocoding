package main

import (
	"fmt"
)

func main() {
	d := [][]int{[]int{-2,-3,3},[]int{-5,-10,1},[]int{10,30,-5}}
	// d = [][]int{[]int{1, -3, 3},[]int{0, -2, 0},[]int{-3, -3, -3}}
	fmt.Println(calculateMinimumHP(d))
}

func calculateMinimumHP(dungeon [][]int) int {
	left := make([][]int, len(dungeon))
	for i := 0; i < len(left); i++ {
		left[i] = make([]int, len(dungeon[i]))
	}
	last := 1 - dungeon[len(dungeon)-1][len(dungeon[0])-1]
	if last < 1 {
		last = 1
	}
	left[len(dungeon)-1][len(dungeon[0])-1] = last
	for i := len(left) - 1; i >= 0; i-- {
		for j := len(left[i]) - 1; j >= 0; j-- {
			if i == len(left) - 1 && j == len(left[i]) - 1 {
				continue
			} else if i == len(left) - 1 {
				left[i][j] = max(1, left[i][j+1]) - dungeon[i][j]
			} else if j == len(left[i]) - 1 {
				left[i][j] = max(1, left[i+1][j]) - dungeon[i][j]
			} else {
				left[i][j] = max(1, min(left[i+1][j], left[i][j+1])) - dungeon[i][j]
			}
		}
	}
	return left[0][0]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
