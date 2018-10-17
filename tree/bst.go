// 二叉查找树
package main

import "fmt"

type BST struct {
	root *Node
	size int
}

type Node struct {
	key   int
	val   interface{}
	left  *Node
	right *Node
}

func newNode(key int, val interface{}) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

func (b *BST) Put(key int, val interface{}) {
	if b.root == nil {
		b.root = newNode(key, val)
		b.size++
		return
	}
	if b.root.put(key, val) {
		b.size++
	}
}

func (b *BST) Get(key int) interface{} {
	if b.root == nil {
		return nil
	}
	return b.root.get(key)
}

func (b *BST) Delete(key int) bool {
	if b.root == nil {
		return false
	}
	var ok bool
	b.root, ok = delNode(b.root, key)
	if ok {
		b.size--
	}
	return ok
}

func (b *BST) Size() int {
	return b.size
}

func delNode(n *Node, key int) (*Node, bool) {
	ok := false
	if n.key > key {
		if n.left != nil {
			n.left, ok = delNode(n.left, key)
		}
		return n, ok
	}
	if n.key < key {
		if n.right != nil {
			n.right, ok = delNode(n.right, key)
		}
		return n, ok
	}
	ok = true
	if n.right == nil {
		return n.left, ok
	}
	t := n
	n = getMinNode(t.right)
	n.right = deleteMin(t.right)
	n.left = t.left
	return n, ok
}

func deleteMin(n *Node) *Node {
	if n.left == nil {
		t := n.right
		n = nil
		return t
	} else {
		n.left = deleteMin(n.left)
	}
	return n
}
func getMinNode(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.left == nil {
		return n
	}
	return getMinNode(n.left)
}

func (n *Node) put(key int, val interface{}) bool {
	if key == n.key {
		n.val = val
		return false
	}
	if key < n.key {
		if n.left == nil {
			n.left = newNode(key, val)
		} else {
			return n.left.put(key, val)
		}
	} else {
		if n.right == nil {
			n.right = newNode(key, val)
		} else {
			return n.right.put(key, val)
		}
	}
	return true
}

func (n *Node) get(key int) interface{} {
	if n.key == key {
		return n.val
	}
	if n.key > key && n.left != nil {
		return n.left.get(key)
	}
	if n.key < key && n.right != nil {
		return n.right.get(key)
	}
	return nil
}

func (b *BST) PreOrder() {
	if b.root == nil {
		return
	}
	b.root.preOrder()
	fmt.Println("-----------------")
}

func (n *Node) preOrder() {
	fmt.Println(n.key, n.val)
	if n.left != nil {
		n.left.preOrder()
	}
	if n.right != nil {
		n.right.preOrder()
	}
}
