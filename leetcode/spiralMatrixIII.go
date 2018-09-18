package main

import "fmt"

func main() {
	fmt.Println(spiralMatrixIII(5, 5, 0, 0))
}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	ret := [][]int{}
	ret = append(ret, []int{r0, c0})
	loop := 1
	all := R * C
	add := [][]int{}
	for len(ret) <= all {
		if loop%2 == 1 {
			add = [][]int{[]int{0, 1}, []int{1, 0}}
		} else {
			add = [][]int{[]int{0, -1}, []int{-1, 0}}
		}
		for i := 0; i < 2; i++ {
			for j := 0; j < loop; j++ {
				r0 += add[i][0]
				c0 += add[i][1]
				if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {
					ret = append(ret, []int{r0, c0})
				}
				if len(ret) >= all {
					return ret
				}
			}
		}
		loop++
	}
	return ret
}
