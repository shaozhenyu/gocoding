package main

import (
	"fmt"
)

func main() {
	fmt.Println(videoStitching([][]int{[]int{0,1}, []int{1,2}}, 5))
}

func videoStitching(clips [][]int, T int) int {
	m := [101]map[int]struct{}{}
	for i := 0; i <= 100; i++ {
		m[i] = make(map[int]struct{})
	}
	for i := 0; i < len(clips); i++ {
		for start := clips[i][0]; start <= clips[i][1]; start++ {
			m[start][i] = struct{}{}
		}
	}
	count := 0
	for i := 0; i <= T; {
		if len(m[i]) == 0 {
			return -1
		}
		end := -1
		for k := range m[i] {
			if clips[k][1] > end {
				end = clips[k][1]
			}
		}
		count++
		i = end + 1
	}
	return count
}
