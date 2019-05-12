package main

import "fmt"

func main() {
	A := []int{2,5,1,2,5}
	B := []int{10,5,2,1,5,2}
	A = []int{1}
	B = []int{1, 3}
	A = []int{1,1,3,5,3,3,5,5,1,1}
	B = []int{2,3,2,1,3,5,3,2,2,1}
	fmt.Println(maxUncrossedLines(A, B))
}

type Node struct {
	x, y int
	count int
}

func maxUncrossedLines(A []int, B []int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
		
		}
	}
}
func maxUncrossedLines(A []int, B []int) int {
	node := make([]Node, 0)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				node = append(node, Node{i, j, 0})
			}
		}
	}

	max, maxIdx, de := -1, -1, -1
	for i := 0; i < len(node); i++ {
		for j := 0; j < len(node); j++ {
			if i == j {
				continue
			}
			if (node[j].x < node[i].x && node[j].y >= node[i].y) || (node[j].x > node[i].x && node[j].y <= node[i].y) {
				node[i].count++
				if node[i].count > max || (node[i].count == max && abs(node[i].x - node[i].y) > de) {
					max = node[i].count
					maxIdx = i
					de = abs(node[i].x - node[i].y)
				}
			}
		}
	}
	if maxIdx == -1 {
		return len(node)
	}
	idx := maxIdx
	for len(node) > 0 {
		fmt.Println(node, node[idx])
		count := 0
		dx, dy := node[idx].x, node[idx].y
		node = append(node[:idx], node[idx+1:]...)
		max, maxIdx, de = -1, -1, -1
		for i := 0; i < len(node); i++ {
			// fmt.Println(dx, dy, node[i])
			if (dx < node[i].x && dy >= node[i].y) || (dx > node[i].x && dy <= node[i].y) {
				// fmt.Println("ddd")
				node[i].count--
			}
			if node[i].count > max || (node[i].count == max && abs(node[i].x - node[i].y) > de) {
				max = node[i].count
				maxIdx = i
				de = abs(node[i].x - node[i].y)
			}
			if node[i].count == 0 {
				count++
			}
		}
		if count == len(node) {
			return count
		}
		idx = maxIdx
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
