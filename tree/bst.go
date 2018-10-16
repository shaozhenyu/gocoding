// 二叉查找树
package main

import "fmt"

type BST struct {
	root *Node
	size int
}

type Node struct {
	Key   int
	Val   interface{}
	Left  *Node
	Right *Node
	Size  int
}

func main() {
	t := new(BST)
	t.Put(1, "1aa")
	t.Put(2, "2aa")
	t.Put(-1, "-1aa")
	t.Put(10, "10aa")
	t = deleteMin(t)
	t.preOrder()
	fmt.Println("-----")
	t = deleteMin(t)
	t.preOrder()
	fmt.Println("-----")
	t = deleteMin(t)
	t.preOrder()
	fmt.Println("-----")
	t = deleteMin(t)
	t.preOrder()
}

func newNode(key int, val interface{}) *Node {
	return &Node{
		Key:  key,
		Val:  val,
		Size: 1,
	}
}

func (b *BST) Put(key int, val interface{}) {
	if n.Size == 0 {
		n.Key = key
		n.Val = val
		n.Size = 1
		return
	}
	if n.Key == key {
		n.Val = val
	}
	if n.Key > key {
		if n.Left != nil {
			n.Left.Put(key, val)
		} else {
			n.Left = initNode(key, val)
		}
	}
	if n.Key < key {
		if n.Right != nil {
			n.Right.Put(key, val)
		} else {
			n.Right = initNode(key, val)
		}
	}
	n.Size = size(n.Left) + size(n.Right) + 1
}

func (n *Node) Get(key int) interface{} {
	if n.Size == 0 {
		return nil
	}
	if n.Key == key {
		return n.Val
	}
	if n.Key > key && n.Left != nil {
		return n.Left.Get(key)
	}
	if n.Key < key && n.Right != nil {
		return n.Right.Get(key)
	}
	return nil
}

func size(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Size
}

func (n *Node) preOrder() {
	if n.Left != nil {
		n.Left.preOrder()
	}
	if n != nil {
		fmt.Println(n.Key, n.Val, n.Size)
	}
	if n.Right != nil {
		n.Right.preOrder()
	}
}

func (n *Node) init() {
	n.Size = 0
	n.Key = 0
	n.Val = nil
	n.Left = nil
	n.Right = nil
}

func deleteMin(n *Node) *Node {
	if n.Size == 0 {
		return n
	}
	if n.Left == nil {
		if n.Right == nil {
			n.init()
			return n
		}
		return n.Right
	} else {
		n.Left = deleteMin(n.Left)
		n.Size--
	}
	return n
}
