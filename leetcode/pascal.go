package main

import "fmt"

func main() {
	fmt.Println(getRow(5))
}

func generate(numRows int) [][]int {
	p := [][]int{}

	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i{
				tmp[j] = 1
			} else {
				tmp[j] = p[i-1][j-1] + p[i-1][j]
			}
		}
		p = append(p, tmp)
	}

	return p
}

func getRow(rowIndex int) []int {
	p := [][]int{}

	for i := 0; i <= rowIndex; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i{
				tmp[j] = 1
			} else {
				fmt.Println(i, j)
				fmt.Println(p[0])
				return p[0]
				tmp[j] = p[i-1][j-1] + p[i-1][j]
			}
		}
		fmt.Println(tmp)
		if i == rowIndex {
			return tmp
		}
		p = append(p, tmp)
	}

	return p[0]
}
