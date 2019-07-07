package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	idx := len(preorder) - 1
	for idx > 0 && preorder[idx] > preorder[0] {
		idx--
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1 : idx+1]),
		Right: bstFromPreorder(preorder[idx+1:]),
	}
}
