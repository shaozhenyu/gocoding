package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		s := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				s[j] = 1
			} else if i >= 3 {
				s[j] = ret[i-2][j-1] + ret[i-2][j]
			}
		}
		ret[i-1] = s
	}
	return ret
}
