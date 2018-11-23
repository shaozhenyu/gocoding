package main

import (
	"fmt"
	"container/list"
)

func main() {
	// matrix := [][]int{[]int{1,2,3}, []int{4,5,6}, []int{7,8,9}}
	matrix := [][]int{[]int{3}, []int{2}}
	fmt.Println(findDiagonalOrder(matrix))
}

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	l := list.New()
	l.PushBack(matrix[0][0])
	i, j := 0, 1
	if j == len(matrix[0]) {
		j = 0
		i = 1
	}
	flag := true
	var end *list.Element
	for ; j < len(matrix[0]) && i < len(matrix); {
		if !flag {
			end = l.Back()
		}
		x, y := i, j
		for x < len(matrix) && y >= 0 {
			if flag {
				l.PushBack(matrix[x][y])
			} else {
				l.InsertAfter(matrix[x][y], end)
			}
			x++
			y--
		}
		flag = !flag
		if j < len(matrix[0]) - 1 {
			j++
		} else {
			i++
		}
	}
	ret := make([]int, len(matrix) * len(matrix[0]))
	idx := 0
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		ret[idx] = v
		idx++
	}
	return ret
}
