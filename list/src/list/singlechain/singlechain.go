//单链表－ 线性表
package singlechain

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func (h *Node) GetHead() *Node {
	if h.Next == nil {
		return nil
	}
	return h.Next
}

func (h *Node) GetLast() *Node {
	if h.Next == nil {
		return nil
	}
	p := h
	for p.Next != nil {
		p = p.Next
		if p.Next == nil {
			return p
		}
	}
	return nil
}

func (h *Node) Insert(d *Node, place int) bool {
	if h.Next == nil {
		if place == 1 {
			h.Next = d
			return true
		} else {
			return false
		}
	}

	p := h
	i := 0
	for p.Next != nil {
		i += 1
		if place == i {
			d.Next = p.Next
			p.Next = d.Next
			return true
		}
		p = p.Next
	}

	if place >= i {
		p.Next = d
		return true
	}
	return false
}

func (h *Node) InsertOnLast(d *Node) bool {
	if h.Next == nil {
		h.Next = d
		return true
	}

	p := h.Next
	for p.Next != nil {
		p = p.Next
	}
	p.Next = d
	return true

}

func (h *Node) List() {
	if h.Next == nil {
		return
	}
	p := h.Next
	for p != nil {
		fmt.Println(p.Data)
		p = p.Next
	}
}

func (h *Node) GetLen() int {
	if h.Next == nil {
		return 0
	}

	i := 0
	p := h.Next
	for p != nil {
		i += 1
		p = p.Next
	}
	return i
}
