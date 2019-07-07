package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{1, &ListNode{1, nil}}}
	fmt.Println(nextLargerNodes(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	val int
	idx int
}

func nextLargerNodes(head *ListNode) []int {
	t := head
	nums := make([]int, 0, 10000)
	for t != nil {
		nums = append(nums, t.Val)
		t = t.Next
	}
	ret := make([]int, len(nums))
	stack := make([]Node, len(nums))
	idx := -1
	for i := 0; i < len(nums); i++ {
		fmt.Println(stack)
		for idx >= 0 {
			if nums[i] > stack[idx].val {
				ret[stack[idx].idx] = nums[i]
				idx--
			} else {
				break
			}
		}
		idx++
		stack[idx] = Node{nums[i], i}
	}
	fmt.Println(stack)
	for idx >= 0 {
		ret[stack[idx].idx] = 0
		idx--
	}
	return ret
}
