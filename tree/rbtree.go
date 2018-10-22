// 红黑树
package main

import "fmt"

var (
	RED   = true
	BLACK = false
)

type RBTree struct {
	root *Node
	size int
}

type Node struct {
	key   int
	val   interface{}
	color bool
	left  *Node
	right *Node
}

func main() {
	rbt := RBTree{}
	rbt.Put(1, 11)
	rbt.Put(2, 22)
	rbt.Put(3, 33)
	rbt.Put(4, 44)
	rbt.Put(5, 55)
	rbt.Put(6, 66)
	rbt.Put(7, 77)
	rbt.Put(8, 88)
	rbt.Put(9, 99)
	midOrder(rbt.root)
	fmt.Println("-----")
	preOrder(rbt.root)
}

func (t *RBTree) Put(key int, val interface{}) {
	t.root = put(t.root, key, val)
	t.root.color = BLACK // 红黑树的跟节点总是黑的
}

func newNode(key int, val interface{}, color bool) *Node {
	return &Node{
		key:   key,
		val:   val,
		color: color,
	}
}

func put(root *Node, key int, val interface{}) *Node {
	if root == nil {
		return newNode(key, val, RED)
	}
	if root.key > key {
		root.left = put(root.left, key, val)
	} else if root.key < key {
		root.right = put(root.right, key, val)
	} else {
		root.val = val
		return root
	}
	if isRed(root.left) && isRed(root.right) {
		root.left.color = BLACK
		root.right.color = BLACK
		root.color = RED
	} else if isRed(root.left) && isRed(root.left.left) {
		root = rotateRight(root)
	} else if isRed(root.right) && !isRed(root.left) {
		root = rotateLeft(root)
	}
	return root
}

// 将左边的红节点转移到右边
func rotateRight(root *Node) *Node {
	x := root.left
	root.left = x.right
	x.right = root
	x.color = root.color
	root.color = RED
	return x
}

// 将右边的红节点转移到左边
func rotateLeft(root *Node) *Node {
	x := root.right
	root.right = x.left
	x.left = root
	x.color = root.color
	root.color = RED
	return x
}

func isRed(n *Node) bool {
	if n == nil {
		return BLACK // 空节点总是黑的
	}
	return n.color
}

func preOrder(t *Node) {
	if t == nil {
		return
	}
	fmt.Println(t.key, t.val, t.color)
	if t.left != nil {
		preOrder(t.left)
	}
	if t.right != nil {
		preOrder(t.right)
	}
}

func midOrder(t *Node) {
	if t == nil {
		return
	}
	if t.left != nil {
		midOrder(t.left)
	}
	fmt.Println(t.key, t.val, t.color)
	if t.right != nil {
		midOrder(t.right)
	}
}
