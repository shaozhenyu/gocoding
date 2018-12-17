package main

import "fmt"

func main() {
	l := &ListNode{1, nil}
	l.Next = &ListNode{2, nil}
	l.Next.Next = &ListNode{3, nil}
	l.Next.Next.Next = &ListNode{4, nil}
	l.Next.Next.Next.Next = &ListNode{5, nil}
	removeNthFromEnd(l, 1)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	t := remove(head, head.Next, n)
	if t >= 0 {
		head = head.Next
	}
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	return head
}

func remove(pre, head *ListNode, n int) int {
	if head != nil {
		t := remove(head, head.Next, n)
		if t == 0 {
			fmt.Println("delete:", head.Val)
			if pre == nil {
				head = head.Next
			} else {
				pre.Next = head.Next
			}
		}
		return t - 1
	} else {
		return n - 1
	}
	return - 1
}
