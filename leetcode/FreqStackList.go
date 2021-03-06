package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

func main() {
	fs := Constructor()
	fs.Push(1)
	fs.Push(2)
	fs.Push(3)
	fs.Push(3)
	fs.Push(2)
	fs.Push(2)
	fs.Push(4)
	fs.Push(4)
	for fs.pq.Len() > 0 {
		fmt.Println(fs.Pop())
	}
}

// list to simulation true heap
type FreqStack struct {
	id    int
	pq    PriorityQueue
	items map[int]*Item
}

func Constructor() FreqStack {
	f := FreqStack{
		pq:    make(PriorityQueue, 0),
		items: make(map[int]*Item),
	}
	heap.Init(&f.pq)
	return f
}

func (this *FreqStack) Push(x int) {
	if _, ok := this.items[x]; ok {
		this.items[x].weight.PushBack(this.id)
		this.items[x].times += 1
		this.items[x].lastIdx = this.id
		this.pq.update(this.items[x])
	} else {
		item := &Item{
			weight: list.New(),
			lastIdx: this.id,
			times:  1,
			num:    x,
			index:  this.id,
		}
		item.weight.PushBack(this.id)
		this.items[x] = item
		heap.Push(&this.pq, item)
	}
	this.id++
}

func (this *FreqStack) Pop() int {
	if this.pq.Len() <= 0 {
		return 0
	}
	item := heap.Pop(&this.pq).(*Item)
	x := item.num
	item.times -= 1
	item.weight.Remove(item.weight.Back())
	if item.times == 0 {
		delete(this.items, x)
	} else {
		item.lastIdx = item.weight.Back().Value.(int)
		heap.Push(&this.pq, item)
		this.items[x] = item
	}
	return x
}

type Item struct {
	weight *list.List
	lastIdx int
	times  int
	num    int
	index  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].times == pq[j].times {
		return pq[i].lastIdx > pq[j].lastIdx
	}
	return pq[i].times > pq[j].times
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item) {
	heap.Fix(pq, item.index)
}
