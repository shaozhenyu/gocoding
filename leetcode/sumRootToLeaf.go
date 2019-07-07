package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{1, nil, nil}
	fmt.Println(sumRootToLeaf(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum *int

func sumRootToLeaf(root *TreeNode) int {
	sum = new(int)
	srtl(root, 0)
	return *sum
}

func srtl(root *TreeNode, pre int) {
	if root == nil {
		return
	}
	pre = pre << 1 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum = (*sum + pre)%1000000007
	} else {
		srtl(root.Left, pre)
		srtl(root.Right, pre)
	}
	return
}
