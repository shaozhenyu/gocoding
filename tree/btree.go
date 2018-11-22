package main

import "fmt"

func main() {
	bTree := newBTree(M)
	for i := 1; i <= 100; i++ {
		bTree.Insert(i)
	}
	fmt.Println("--------")
	bTree.Traverse()
	for i := 1; i <= 1; i++ {
		fmt.Println(i, bTree.Delete(i))
	}
	fmt.Println("--------")
	bTree.Traverse()
}

const M int = 5

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

func (b *BTree) Delete(key int) bool {
	find, i, node := b.Search(key)
	if find {
		b.Root = node.delete(key, i)
	}
	return find
}

func (n *BNode) delete(key int, idx int) *BNode {
	// 非叶子节点
	if n.Child[idx] != nil {
		child := n.Child[idx]
		n.Key[idx] = child.Key[1]
		return n.Child[idx].delete(child.Key[1], 1)
	} else {
		for i := idx; i < n.Num; i++ {
			n.Key[i] = n.Key[i+1]
			// n.Child[i] = n.Child[i+1] // do not need
		}
		n.Num--
		for n.Num < (M-1)/2 && n.Parent != nil {
			ok, n := n.restore()
			if !ok {
				n = n.mergeNode()
			}
		}
	}
	for n.Parent != nil && n.Parent.Num > 0 {
		n = n.Parent
	}
	return n
}

func (n *BNode) restore() (bool, *BNode) {
	if n.Parent == nil {
		return false, nil
	}
	parent := n.Parent
	i := 0
	for ; parent.Child[i] != n; i++ {
	}
	// n 有左兄弟节点
	if i > 0 {
		b := parent.Child[i-1]
		if b.Num > (M-1)/2 {
			// 将parent节点下移
			for j := n.Num; j >= 0; j-- {
				n.Key[j+1] = n.Key[j]
			}
			n.Key[1] = parent.Key[i]
			parent.Key[i] = b.Key[b.Num]
			n.Num++
			b.Num--
			return true, parent
		}
	}

	// n 有右兄弟节点
	if i < parent.Num {
		b := parent.Child[i+1]
		if b.Num > (M-1)/2 {
			n.Key[n.Num+1] = parent.Key[i+1]
			n.Num++
			parent.Key[i+1] = b.Key[1]
			for j := 1; j < b.Num; j++ {
				b.Key[j] = b.Key[j+1]
			}
			b.Num--
			return true, parent
		}
	}
	return false, n
}

func (n *BNode) mergeNode() *BNode {
	parent := n.Parent
	i := 0
	for ; parent.Child[i] != n; i++ {
	}
	// 存在左兄弟节点
	if i > 0 {
		b := parent.Child[i-1]
		b.Num++
		b.Key[b.Num] = parent.Key[i]
		for j := 1; j < n.Num; j++ {
			b.Num++
			b.Key[b.Num] = n.Key[j]
		}
		for j := i - 1; j < parent.Num; j++ {
			parent.Key[j] = parent.Key[j+1]
			parent.Child[j] = parent.Child[j+1]
		}
		parent.Num--
		parent.Child[i-1] = b

		// // 检查parent是否合法
		// if parent.Num < (M-1)/2 && parent.Parent != nil {
		// 	ok, parent := parent.restore()
		// 	if !ok {
		// 		parent = parent.mergeNode()
		// 	}
		// }
	} else {
		b := parent.Child[i+1]
		n.Num++
		n.Key[n.Num] = parent.Key[1]
		for j := 1; j <= b.Num; j++ {
			n.Num++
			n.Key[n.Num] = b.Key[j]
		}
		parent.Num--
		for j := 1; j <= parent.Num; j++ {
			parent.Key[j] = parent.Key[j+1]
			parent.Child[j] = parent.Child[j+1]
		}

		// n = parent
		// // 检查parent是否合法
		// if parent.Num < (M-1)/2 && parent.Parent != nil {
		// 	ok, parent := parent.restore()
		// 	if !ok {
		// 		parent = parent.mergeNode()
		// 	}
		// }
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
