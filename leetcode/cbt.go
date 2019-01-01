package main

import "fmt"

func main() {
	c := Constructor(nil)
	fmt.Println(c.Insert(1))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
	add []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	add := make([]*TreeNode, 0, 1000)
	if root == nil {
		return CBTInserter{root, add}
	}
	add = append(add, root)
	for {
		t := add[0]
		if t.Left == nil {
			break
		}
		add = append(add ,t.Left)
		if t.Right == nil {
			break
		}
		add = append(add, t.Right)
		add = add[1:]
	}
	return CBTInserter{root, add}
}


func (this *CBTInserter) Insert(v int) int {
	if len(this.add) == 0 {
		this.root = &TreeNode{v, nil, nil}
		this.add = append(this.add, this.root)
		return 0
	}
	t := this.add[0]
	if t.Left == nil {
		t.Left = &TreeNode{v, nil, nil}
		this.add = append(this.add, t.Left)
	} else if t.Right == nil {
		t.Right = &TreeNode{v, nil, nil}
		this.add = append(this.add, t.Right)
		this.add = this.add[1:]
	}
	return t.Val
}


func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
