package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var sum int

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	btg(root, &sum)
	return root
}

func btg(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	btg(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	btg(root.Left, sum)
}
