package main

import "fmt"

func main() {
	fmt.Println(maxAncestorDiff(&TreeNode{1, nil, &TreeNode{2, nil, nil}}))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	max := -2 << 31
	mad(root, &max)
	return max
}

func mad(root *TreeNode, max *int) (m map[int]bool) {
	m = make(map[int]bool)
	if root == nil {
		return
	}
	left := mad(root.Left, max)
	right := mad(root.Right, max)
	for key := range left {
		if abs(root.Val - key) > *max {
			*max = abs(root.Val - key)
		}
		m[key] = true
	}
	for key := range right {
		if abs(root.Val - key) > *max {
			*max = abs(root.Val - key)
		}
		m[key] = true
	}
	m[root.Val] = true
	return
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
