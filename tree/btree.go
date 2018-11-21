package main

import "fmt"

func main() {
	bTree := newBTree(M)
	for i := 1; i <= 20; i++ {
		bTree.Insert(i)
	}
	bTree.Traverse()
}

const M int = 4

type BTree struct {
	M    int // 阶
	Root *BNode
}

type BNode struct {
	Num    int
	Key    [M + 1]int
	Parent *BNode
	Child  [M + 1]*BNode
}

func newBTree(m int) *BTree {
	return &BTree{
		M: m,
	}
}

func newBNode(key int) *BNode {
	b := &BNode{
		Num: 1,
	}
	b.Key[1] = key
	return b
}

func (b *BTree) Search(key int) (find bool, idx int, node *BNode) {
	if b.Root == nil {
		return
	}
	t := b.Root
	for t != nil {
		idx = t.Num
		for ; idx > 0 && key <= t.Key[idx]; idx-- {
			if t.Key[idx] == key {
				find, node = true, t
				return
			}
		}
		if t.Child[idx] == nil {
			node = t
		}
		t = t.Child[idx]
	}
	return
}

func (b *BTree) Insert(key int) bool {
	if b.Root == nil {
		b.Root = newBNode(key)
		return true
	}
	find, _, node := b.Search(key)
	if !find {
		var i = node.Num
		for ; i > 0 && node.Key[i] > key; i-- {
			node.Key[i+1] = node.Key[i]
		}
		node.Key[i+1] = key
		node.Num++
		if node.Num < b.M {
			return true
		} else {
			parent := node.Split()
			for parent.Parent != nil {
				parent = parent.Parent
			}
			b.Root = parent
			return true
		}
	}
	return false
}

func (n *BNode) Split() *BNode {
	parent := n.Parent
	if parent == nil {
		parent = &BNode{}
	}
	// 拆分成两个新的node
	mid := n.Num/2 + 1
	newNode := &BNode{}
	newNode.Num = M - mid
	for i := 1; i <= newNode.Num; i++ {
		newNode.Key[i] = n.Key[mid+i]
		newNode.Child[i-1] = n.Child[mid+i-1]
		if newNode.Child[i-1] != nil {
			newNode.Child[i-1].Parent = newNode // !!
		}
		n.Key[mid+1] = 0
		// n.Child[mid+i-1] = nil
	}
	newNode.Child[newNode.Num] = n.Child[n.Num]
	if newNode.Child[newNode.Num] != nil {
		newNode.Child[newNode.Num].Parent = newNode // !!
	}
	n.Num = mid - 1
	n.Parent = parent
	newNode.Parent = parent
	// 将中间节点插入到parent
	parentNum := parent.Num
	midNode := n.Key[mid]
	k := parentNum
	for ; midNode < parent.Key[k]; k-- {
		parent.Key[k+1] = parent.Key[k]
		parent.Child[k+1] = parent.Child[k]
	}
	parent.Key[k+1] = midNode
	parent.Child[k] = n
	parent.Child[k+1] = newNode
	parent.Num++
	if parent.Num >= M {
		return parent.Split()
	}
	return parent
}

func (b *BTree) Traverse() {
	queue := make([]*BNode, 0)
	if b.Root != nil {
		queue = append(queue, b.Root)
	}
	idx := 0
	for len(queue) > idx {
		now := queue[idx]
		idx++
		fmt.Print("[")
		for i := 1; i <= now.Num; i++ {
			fmt.Printf(" %d ", now.Key[i])
		}
		fmt.Print("]\n")
		for i := 0; i <= now.Num; i++ {
			if now.Child[i] != nil {
				queue = append(queue, now.Child[i])
			}
		}
	}
}
