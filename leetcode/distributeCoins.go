package main

import "fmt"

func main() {

}

func distributeCoins(root *TreeNode) int {
	ret, _ := dc(root)
	return ret
}

func dc(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftStep, leftGet := dc(root.Left)
	rightStep, rightGet := dc(root.Right)
	get := root.Val + leftGet + rightGet
	return leftStep + rightStep + abs(get), get
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
