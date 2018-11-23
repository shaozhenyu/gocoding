package main

import "fmt"

func main() {
	heights := []int{2,1,5,6,2,3}
	fmt.Println(largestRectangleArea(heights))
}

type Node struct {
	val int
	idx int
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	ret := heights[0]
	stack := make([]Node, len(heights)
	idx := 0
	stack[idx] = Node{heights[0], 0}
	for i := 1; i < len(heights); i++ {
		if idx < 0 {
			idx++
			stack[idx] = Node{heights[i], i}
			if ret < heights[i] {
				ret = heights[i]
			}
		} else {
			for heights[i] < stack[idx] {
				idx--
			}
		}
	}
	return ret
}
