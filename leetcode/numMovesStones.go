package main

import (
	"fmt"
	"sort"
)

func main () {
	fmt.Println(numMovesStones(3, 5, 1))
}

func numMovesStones(a int, b int, c int) []int {
	s := []int{a, b, c}
	sort.Ints(s)
	a, b, c = s[0], s[1], s[2]
	fmt.Println(a, b, c)
	min, max := 0, 0
	if b - a > 1 {
		min++
		max += b - a - 1
	}
	if c - b > 1 {
		min++
		max += c - b - 1
	}
	if b - a == 2 || c - b == 2 {
		min = 1
	}
	return []int{min, max}
}
