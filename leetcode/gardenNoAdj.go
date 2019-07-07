package main

import "fmt"

func main() {
	N := 3
	paths := [][]int{[]int{1,2},[]int{2,3},[]int{3,1}}
	fmt.Println(gardenNoAdj(N, paths))
}

func gardenNoAdj(N int, paths [][]int) []int {
    m := make(map[int]map[int]bool)
	for i := 1; i <= N; i++ {
		m[i] = make(map[int]bool)
	}
	for i := 0; i < len(paths); i++ {
		m[paths[i][0]][paths[i][1]] = true
		m[paths[i][1]][paths[i][0]] = true
	}
	ans := make([]int, N+1)
	for i := 1; i <= N; i++ {
		t := [5]bool{}
		for key := range m[i] {
			t[key] = true
		}
		for j := 1; j <= 4; j++ {
			if !t[j] {
				ans[i] = j
				break
			}
		}
	}
	return ans[1:]
}
