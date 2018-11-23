package main

import "fmt"

func main() {
	matrix := [][]int{[]int{1,2,3}, []int{4,5,6}, []int{7,8,9}}
	matrix = [][]int{[]int{3,0,1,4,2},[]int{5,6,3,2,1},[]int{1,2,0,1,5},[]int{4,1,0,1,7},[]int{1,0,3,0,5}}
	// [2,1,4,3],[]int{1,1,2,2},[]int{1,2,2,4}}
	m := Constructor(matrix)
	fmt.Println(m)
	fmt.Println(m.SumRegion(2, 1, 4, 3))
}

type NumMatrix struct {
	m map[string]int
}


func Constructor(matrix [][]int) NumMatrix {
	ma := make(map[string]int)
	m := len(matrix)
	n := len(matrix[0])
	for a := 0; a < 1; a++ {
		for b := 0; b < 1; b++ {
			for c := a; c < m; c++ {
				for d := b; d < n; d++ {
					vk := fmt.Sprintf("%d+%d+%d+%d", a, b, c, d)
					ma[vk] = matrix[c][d]
					fmt.Println(vk, ma[vk])
					if c > 0 {
						ma[vk] += ma[fmt.Sprintf("%d+%d+%d+%d", a, b, c-1, d)]
					}
					if d > 0 {
						ma[vk] += ma[fmt.Sprintf("%d+%d+%d+%d", a, b, c, d-1)]
					}
					if c > 0 && d > 0 {
						ma[vk] -= ma[fmt.Sprintf("%d+%d+%d+%d", a, b, c-1, d-1)]
					}
					fmt.Println(ma[vk])
				}
			}
		}
	}
	return NumMatrix{ma}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ret := this.m[fmt.Sprintf("%d+%d+%d+%d", 0, 0, row2, col2)]
	if row1 == 0 && col1 == 0 {
		return ret
	}
	fmt.Println("1:", ret)
	if row1 > 0 {
		ret -= this.m[fmt.Sprintf("%d+%d+%d+%d", 0, 0, row1-1, col2)]
		fmt.Println("2: ", ret)
	}
	if col1 > 0 {
		ret -= this.m[fmt.Sprintf("%d+%d+%d+%d", 0, 0, row2, col1-1)]
		fmt.Println("3: ", ret)
	}
	if row1 > 0 && col1 > 0 {
		ret += this.m[fmt.Sprintf("%d+%d+%d+%d", 0, 0, row1-1, col1-1)]
	}
	return ret
}
