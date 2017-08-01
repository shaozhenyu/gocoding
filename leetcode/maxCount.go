package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxCount(39999, 39999, [][]int{[]int{19999, 19999}}))
}

func maxCount(m int, n int, ops [][]int) int {
	shortM := m
	shortN := n
	for i := 0; i < len(ops); i++ {
		if shortM > ops[i][0] {
			shortM = ops[i][0]
		}
		if shortN > ops[i][1] {
			shortN = ops[i][1]
		}
	}
	return shortM * shortN
}
