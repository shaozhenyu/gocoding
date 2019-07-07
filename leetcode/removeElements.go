package main

import "fmt"

func main() {
	fmt.Println(removeElements(&ListNode{1, nil}, 1))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	pre := head
	t := head.Next
	for t != nil {
		if t.Val == val {
			pre.Next = t.Next
			t = pre.Next
		} else {
			pre = t
			t = t.Next
		}
	}
	return head
}
