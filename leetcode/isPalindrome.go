package main

import "fmt"

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{2, nil}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{1, nil}
	fmt.Println(isPalindrome(head))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	count := 0
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		count++
	}
	for fast.Next != nil {
		fast = fast.Next
		count++
	}
	t := slow.Next
	size := count
	for count > 1 && t != nil {
		count--
		slow.Next = t.Next
		t.Next = fast.Next
		fast.Next = t
		t = slow.Next
	}
	slow = slow.Next
	for size > 0 {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
		size--
	}
	return true
}
