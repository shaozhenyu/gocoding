package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var xLevel, xFather, yLevel, yFather int
	traver(root, x, y, 0, 0, &xLevel, &xFather, &yLevel, &yFather)
	if xLevel != yLevel || xLevel == 0 {
		return false
	}
	if xFather == yFather || xFather == 0 {
		return false
	}
	return true
}

func traver(root *TreeNode, x, y, f, level int, xLevel, xFather, yLevel, yFather *int) {
	if root == nil {
		return
	}
	if root.Val == x {
		*xLevel = level
		*xFather = f
		return
	}
	if root.Val == y {
		*yLevel = level
		*yFather = f
		return
	}
	traver(root.Left, x, y, root.Val, level+1, xLevel, xFather, yLevel, yFather)
	traver(root.Right, x, y, root.Val, level+1, xLevel, xFather, yLevel, yFather)
}


