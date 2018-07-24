package main

import (
	"fmt"
	"container/list"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}
	root.Right = &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}
	fmt.Println(postorderTraversal(root))
}

// 迭代算法
func postorderTraversal(root *TreeNode) []int {
	l := list.New()
	l.PushBack(root)
	ret := []int{}
	for l.Len() > 0 {
		last := l.Back()
		lastNode, _ := last.Value.(*TreeNode)
		l.Remove(last)
		if lastNode != nil {
			ret = append([]int{lastNode.Val}, ret...)
		}
		if lastNode.Left != nil {
			l.PushBack(lastNode.Left)
		}
		if lastNode.Right != nil {
			l.PushBack(lastNode.Right)
		}
	}
	return ret
}

// 递归算法
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	ret = append(ret, postorderTraversal(root.Left)...)
	ret = append(ret, postorderTraversal(root.Right)...)
	ret = append(ret, root.Val)
	return ret
}

