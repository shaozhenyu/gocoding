package main

import "fmt"

func main() {
	list := Constructor()
	list.AddAtHead(2)
	list.AddAtHead(1)
	list.AddAtIndex(2, 7)
	list.print()
	list.AddAtTail(3)
	list.AddAtTail(4)
	list.AddAtIndex(4, 5)
	list.AddAtIndex(0, 0)
	list.AddAtIndex(1, 10)
	list.DeleteAtIndex(5)
	list.print()
	fmt.Println(list.Get(5))

}

func (l *MyLinkedList) print() {
	fmt.Println("size: ", l.Size)
	t := l.Head
	for t != nil {
		fmt.Printf("%d ", t.Val)
		t = t.Next
	}
	fmt.Println()
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Val int
	Next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size {
		return -1
	}
	if index == this.Size - 1 {
		return this.Tail.Val
	}
	t := this.Head
	for index > 0 {
		t = t.Next
		index--
	}
	return t.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	t := &Node{val, this.Head}
	this.Head = t
	if this.Tail == nil {
		this.Tail = this.Head
	}
	this.Size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	t := &Node{val, nil}
	if this.Tail == nil {
		this.Tail = t
		this.Head = t
	} else {
		this.Tail.Next = t
		this.Tail = t
	}
	this.Size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	pre := this.Head
	for index > 1 {
		index--
		pre = pre.Next
	}
	t := &Node{val, nil}
	t.Next = pre.Next
	pre.Next = t
	this.Size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size {
		return
	}
	this.Size--
	if index == 0 {
		this.Head = this.Head.Next
		if this.Head == nil {
			this.Tail = nil
		}
		return
	}
	pre := this.Head
	t := pre.Next
	for index > 1 {
		pre = pre.Next
		t = t.Next
		index--
	}
	pre.Next = t.Next
	if t == this.Tail {
		this.Tail = pre
	}
	t = nil
}
