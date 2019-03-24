package main

import "fmt"

func main() {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var alp = []byte{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func smallestFromLeaf(root *TreeNode) string {
	return sf(root, "")
}

func sf(root *TreeNode, suffix string) string {
	if root == nil {
		return suffix
	}
	suffix = string(alp[root.Val]) + suffix
	left := sf(root.Left, suffix)
	right := sf(root.Right, suffix)
	return min(left, right)
}

func min(x, y string) string {
	if x < y {
		return x
	}
	return y
}
