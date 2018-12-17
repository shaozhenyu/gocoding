package main

import "fmt"

func main() {
	nums := []int{5, 1, 6, 2, 1, 3}
	fmt.Println(countSmaller(nums))
}

type Node struct {
	key   int
	count int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func countSmaller(nums []int) []int {
	ret := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] = 
	}
	return ret
}

func newNode(key int, val interface{}) *Node {
	return &Node{
		key: key,
		count: 1,
	}
}

func (b *BST) Get(key int) int {
	if b.root == nil {
		return 0
	}
	return b.root.get(key)
}

func (n *Node) get(key int) int {
	if n.key == key {
		return n.count
	}
	if n.key > key && n.left != nil {
		return n.left.get(key)
	}
	if n.key < key && n.right != nil {
		return n.right.get(key)
	}
	return nil
}

func (b *BST) Put(key int, val interface{}) {
	if b.root == nil {
		b.root = newNode(key, val)
		return
	}
	b.root.put(key, val)
}

func (n *Node) put(key int) {
	if key == n.key {
		return
	}
	if key < n.key {
		n.count++
		if n.left == nil {
			n.left = newNode(key, val)
		} else {
			n.left.put(key, val)
		}
	} else {
		if n.right == nil {
			n.right = newNode(key, val)
		} else {
			n.right.put(key, val)
		}
	}
	return
}
