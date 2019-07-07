package main

import (
	"fmt"
	"strconv"
)

func main() {
	t := recoverFromPreorder("1-2-3--4---5--6")
	tprint(t)
}

func tprint(t *TreeNode) {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	tprint(t.Left)
	tprint(t.Right)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {
	fmt.Println("xxx1, ", S)
	if S == "" {
		return nil
	}
	idx := 0
	for idx < len(S) && S[idx] != '-' {
		idx++
	}
	val, _ := strconv.Atoi(S[:idx])
	root := &TreeNode{val, nil, nil}
	rfp(S[idx:], root, 0, true)
	return root
}

func rfp(S string, root *TreeNode, level int, isLeft bool) string {
	fmt.Println("xxx1, ", S)
	if S == "" {
		return ""
	}
	idx := 0
	for S[idx] == '-' {
		idx++
	}
	nowLevel := idx
	if nowLevel <= level {
		return S
	}
	S = S[idx:]
	idx = 0
	for idx < len(S) && S[idx] != '-' {
		idx++
	}
	val, _ := strconv.Atoi(S[:idx])
	S = S[idx:]
	root.Left = &TreeNode{val, nil, nil}
	S = rfp(S, root.Left, nowLevel, true)
	if S == "" {
		return S
	}
	idx = 0
	for S[idx] == '-' {
		idx++
	}
	nowLevel = idx
	if nowLevel <= level {
		return S
	}
	S = S[idx:]
	idx = 0
	for idx < len(S) && S[idx] != '-' {
		idx++
	}
	val, _ = strconv.Atoi(S[:idx])
	S = S[idx:]
	root.Right = &TreeNode{val, nil, nil}
	S = rfp(S, root.Right, nowLevel, true)
	return S
}
