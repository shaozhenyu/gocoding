package main

import "fmt"

func main() {
	blocked := [][]int{[]int{0,1},[]int{1,0}}
	source := []int{0,0}
	target := []int{0,2}
	blocked = [][]int{[]int{691938,300406},[]int{710196,624190},[]int{858790,609485},[]int{268029,225806},[]int{200010,188664},[]int{132599,612099},[]int{329444,633495},[]int{196657,757958},[]int{628509,883388}}
	source = []int{655988,180910}
	target = []int{267728,840949}
	fmt.Println(isEscapePossible(blocked, source, target))
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(blocked) == 0 {
		return true
	}
	bm := make(map[[2]int]struct{})
	for i := 0; i < len(blocked); i++ {
		bm[[2]int{blocked[i][0], blocked[i][1]}] = struct{}{}
	}
	rs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	var travel func(x, y int, visited map[[2]int]struct{}, targetX, targetY int) bool
	travel = func(x, y int, visited map[[2]int]struct{}, targetX, targetY int) bool {
		key := [2]int{x, y}
		if _, ok := visited[key]; ok {
			return false
		}
		if _, ok := bm[key]; ok {
			return false
		}
		visited[key] = struct{}{}
		if x == targetX && y == targetY {
			visited[[2]int{targetX, targetY}] = struct{}{}
		}
		if _, ok := visited[[2]int{targetX, targetY}]; ok {
			return true
		}
		if len(visited) >= 500 {
			return true
		}
		for _, r := range rs {
			i := x + r[0]
			j := y + r[1]
			if i >= 0 && j >= 0 && i <= 1000000 && j <= 1000000 {
				if travel(i, j, visited, targetX, targetY) {
					return true
				}
			}
		}
		return false
	}
	return travel(target[0], target[1], make(map[[2]int]struct{}), source[0], source[1]) && travel(source[0], source[1], make(map[[2]int]struct{}), target[0], target[1])
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
