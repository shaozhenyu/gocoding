package main

import "fmt"

func main() {
	A := [][]int{[]int{0, 1}, []int{1, 0}}
	fmt.Println(shortestBridge(A))
}

type PPoint struct {
	x, y int
}

func shortestBridge(A [][]int) int {
	m1 := make(map[string]PPoint)
	m2 := make(map[string]PPoint)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				if len(m1) == 0 {
					find(i, j, A, m1)
				} else {
					find(i, j, A, m2)
				}
			}
		}
	}
	ret := 10000
	fmt.Println(m1)
	fmt.Println(m2)
	for _, p1 := range m1 {
		for _, p2 := range m2 {
			ret = min(ret, abs(p2.x-p1.x)+abs(p2.y-p1.y)-1)
			if ret == 1 {
				return 1
			}
		}
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		x = x * -1
	}
	return x
}

func find(i, j int, A [][]int, m map[string]PPoint) {
	if A[i][j] == 0 {
		return
	}
	m[fmt.Sprintf("%d+%d", i, j)] = PPoint{i, j}
	A[i][j] = 0
	if i > 0 {
		find(i-1, j, A, m)
	}
	if i < len(A)-1 {
		find(i+1, j, A, m)
	}
	if j > 0 {
		find(i, j-1, A, m)
	}
	if j < len(A[i])-1 {
		find(i, j+1, A, m)
	}
}
