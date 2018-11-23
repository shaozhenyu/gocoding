package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type BNode struct {
	key int
}

func (b BNode) Less(b1 Item) bool {
	return b.key < b1.(BNode).key
}

func main() {
	degree := 2
	b := New(degree)
	for i := 1; i <= 4; i++ {
		bn := BNode{i}
		b.ReplaceOrInsert(bn)
	}
	fmt.Println("insert")
	b.root.print(0)
	for i := 1; i <= 2; i++ {
		bn := BNode{i}
		b.Delete(bn)
	}
	fmt.Println("delete")
	b.root.print(0)
}

const (
	DefaultFreeListSize = 32
)

var (
	nilItems    = make(items, 16)
	nilChildren = make(children, 16)
)

type Item interface {
	Less(than Item) bool
}

type items []Item

func (s items) find(item Item) (index int, found bool) {
	i := sort.Search(len(s), func(i int) bool {
		return item.Less(s[i])
	})
	// s[i-1] equal item find
	if i > 0 && !s[i-1].Less(item) {
		return i - 1, true
	}
	return i, false
}

func (s *items) insertAt(index int, item Item) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = item
}

func (s *items) truncate(index int) {
	var toClear items
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilItems):]
	}
}

func (s *items) removeAt(index int) Item {
	item := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return item
}

type children []*node

func (s *children) insertAt(index int, n *node) {
	*s = append(*s, nil)
	if index > len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = n
}

func (s *children) truncate(index int) {
	var toClear children
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilChildren):]
	}
}

type node struct {
	items    items
	children children
	cow      *copyOnWriteContext
}

func (n *node) mutableFor(cow *copyOnWriteContext) *node {
	if n.cow == cow {
		return n
	}
	out := cow.newNode()
	if cap(out.items) >= len(n.items) {
		out.items = out.items[:len(n.items)]
	} else {
		out.items = make(items, len(n.items), cap(n.items))
	}
	copy(out.items, n.items)
	if cap(out.children) >= len(n.children) {
		out.children = out.children[:len(n.children)]
	} else {
		out.children = make(children, len(n.children), cap(n.children))
	}
	copy(out.children, n.children)
	return out
}

func (n *node) mutableChild(i int) *node {
	c := n.children[i].mutableFor(n.cow)
	n.children[i] = c
	return c
}

func (n *node) maybeSplitChild(i, maxItems int) bool {
	if len(n.children[i].items) < maxItems {
		return false
	}
	first := n.mutableChild(i)
	item, second := first.split(maxItems / 2)
	n.items.insertAt(i, item)
	n.children.insertAt(i+1, second)
	return true
}

func (n *node) insert(item Item, maxItems int) Item {
	i, found := n.items.find(item)
	if found {
		out := n.items[i]
		n.items[i] = item
		return out
	}
	// 叶子结点
	if len(n.children) == 0 {
		n.items.insertAt(i, item)
		return nil
	}
	fmt.Println("insert :", i, item)
	if n.maybeSplitChild(i, maxItems) {
		change := n.items[i]
		switch {
		case item.Less(change):
			// do nothing
		case change.Less(item):
			i++
		default:
			// find
			out := n.items[i]
			n.items[i] = item
			return out
		}
	}
	return n.mutableChild(i).insert(item, maxItems)
}

func (n *node) split(i int) (Item, *node) {
	item := n.items[i]
	next := n.cow.newNode()
	next.items = append(next.items, n.items[i+1:]...)
	n.items.truncate(i)
	if len(n.children) > 0 {
		next.children = append(next.children, n.children[i+1:]...)
		n.children.truncate(i + 1)
	}
	return item, next
}

func (n *node) remove(item Item, minItems int) Item {
	var i int
	var found bool
	i, found = n.items.find(item)
	if len(n.children) == 0 {
		if found {
			return n.items.removeAt(i)
		}
		return nil
	}
	if len(n.children[i].items) <= minItems {
		return n.growChildAndRemove(i, item, minItems)
	}

	return nil
}

func (n *node) growChildAndRemove(i int, item Item, minItems int) Item {
	// 左子节点有多余的item可以使用
	// if i > 0 && len(n.children[i-1].items) > minItems {

	// }
	fmt.Println("growChildAndRemove i:", i)
	return nil
}

type FreeList struct {
	mu       sync.Mutex
	freelist []*node
}

func (f *FreeList) newNode() (n *node) {
	f.mu.Lock()
	defer f.mu.Unlock()

	index := len(f.freelist) - 1
	if index < 0 {
		return new(node)
	}
	n = f.freelist[index]
	f.freelist[index] = nil
	f.freelist = f.freelist[:index]
	return
}

type copyOnWriteContext struct {
	freelist *FreeList
}

func (c *copyOnWriteContext) newNode() (n *node) {
	n = c.freelist.newNode()
	n.cow = c
	return
}

type BTree struct {
	degree int
	length int
	root   *node
	cow    *copyOnWriteContext
}

func (t *BTree) maxItems() int {
	return t.degree*2 - 1
}

func (t *BTree) minItems() int {
	return t.degree - 1
}

func (t *BTree) ReplaceOrInsert(item Item) Item {
	if item == nil {
		panic("nil item being added to BTree")
	}
	if t.root == nil {
		t.root = t.cow.newNode()
		t.root.items = append(t.root.items, item)
		t.length++
		return nil
	} else {
		t.root = t.root.mutableFor(t.cow)
		if len(t.root.items) >= t.maxItems() {
			item2, second := t.root.split(t.maxItems() / 2)
			oldRoot := t.root
			t.root = t.cow.newNode()
			t.root.items = append(t.root.items, item2)
			t.root.children = append(t.root.children, oldRoot, second)
		}
	}
	out := t.root.insert(item, t.maxItems())
	if out == nil {
		t.length++
	}
	return out
}

func (t *BTree) Delete(item Item) Item {
	if t.root == nil || len(t.root.items) == 0 {
		return nil
	}
	t.root = t.root.mutableFor(t.cow)
	out := t.root.remove(item, t.minItems())
	if len(t.root.items) == 0 && len(t.root.children) > 0 {
		// oldRoot := t.root
		t.root = t.root.children[0]
	}
	if out != nil {
		t.length--
	}
	return out
}

// new
func New(degree int) *BTree {
	return NewWithFreeList(degree, NewFreeList(DefaultFreeListSize))
}

func NewWithFreeList(degree int, f *FreeList) *BTree {
	if degree <= 1 {
		panic("bad degree")
	}
	return &BTree{
		degree: degree,
		cow:    &copyOnWriteContext{freelist: f},
	}
}

func NewFreeList(size int) *FreeList {
	return &FreeList{freelist: make([]*node, 0, size)}
}

// test
func (n *node) print(level int) {
	fmt.Printf("%sNODE:%v\n", strings.Repeat("  ", level), n.items)
	for _, c := range n.children {
		c.print(level + 1)
	}
}
