package main

import "fmt"

func main() {
	t := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	recoverTree(t)
	fmt.Println(t.Val, t.Left.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var firstNode, secNode, preNode *TreeNode

func recoverTree(root *TreeNode) {
	firstNode, secNode, preNode = nil, nil, nil
	preNode = &TreeNode{-2 << 31, nil, nil}
	inOrder(root)
	firstNode.Val, secNode.Val = secNode.Val, firstNode.Val
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	if firstNode == nil && preNode.Val > root.Val {
		firstNode = preNode
	}
	if firstNode != nil && preNode.Val > root.Val {
		secNode = root
	}
	preNode = root
	inOrder(root.Right)
}
