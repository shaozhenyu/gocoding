package main

import "fmt"

func main() {
	N := 5
	lamps := [][]int{[]int{0,0}, []int{1, 0}}
	queries := [][]int{[]int{1, 1}, []int{1,1}}
	fmt.Println(gridIllumination(N, lamps, queries))
}

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	x1, x2, x3, x4 := make(map[int]int), make(map[int]int), make(map[int]int), make(map[int]int)
	open := make(map[int]map[int]struct{})
	for i := 0; i < len(lamps); i++ {
		x, y := lamps[i][0], lamps[i][1]
		x1[x]++
		x2[y]++
		x3[x-y]++
		x4[x+y]++
		if _, ok := open[x]; !ok {
			open[x] = make(map[int]struct{})
		}
		open[x][y] = struct{}{}
	}
	answer := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		x, y := queries[i][0], queries[i][1]
		if x1[x] > 0 || x2[y] > 0 || x3[x-y] > 0 || x4[x+y] > 0 {
			answer[i] = 1
		}
		a := x - 1
		if a < 0 {
			a++
		}
		for ; a <= x+1 && a < N; a++ {
			b := y - 1
			if b < 0 {
				b++
			}
			for ; b <= y+1 && b < N; b++ {
				if _, ok := open[a]; ok {
					if _, ok := open[a][b]; ok {
						x1[a]--
						x2[b]--
						x3[a-b]--
						x4[a+b]--
						delete(open[a], b)
					}
				}
			}
		}
	}
	return answer
}

func checkOpen(x, y, N int, open map[int]map[int]struct{}) bool {
	if x >= 0 && x < N && y >= 0 && y < N {
		if _, ok := open[x]; ok {
			if _, ok := open[x][y]; ok {
				return true
			}
		}
	}
	return false
}
