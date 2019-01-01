package main

import (
	"container/list"
	"fmt"
)

// ["AllOne","inc","inc","inc","inc","getMaxKey","inc","inc","inc","dec","inc","inc","inc","getMaxKey"]
// [[],["hello"],["goodbye"],["hello"],["hello"],[],["leet"],["code"],["leet"],["hello"],["leet"],["code"],["code"],[]]

func main() {
	all := Constructor()
	all.Inc("hello")
	all.Inc("goodbye")
	all.Inc("hello")
	all.Inc("hello")
	fmt.Println(all.GetMaxKey())
	all.Inc("leet")
	all.Inc("code")
	all.Inc("leet")
	all.Dec("hello")

	for e := all.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%s,", e.Value.(string))
	}
	fmt.Println()

	all.Inc("leet")
	all.Inc("code")
	all.Inc("code")
	fmt.Println(all.GetMaxKey())

	fmt.Println(all.val)

	for e := all.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(string))
	}
}

type AllOne struct {
	val  map[string]int
	list *list.List
	ele  map[string]*list.Element
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		val:  make(map[string]int),
		list: list.New(),
		ele:  make(map[string]*list.Element),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if _, ok := this.val[key]; !ok {
		this.val[key] = 1
		this.list.PushFront(key)
		this.ele[key] = this.list.Front()
	} else {
		this.val[key]++
		val := this.val[key]
		e := this.ele[key]
		next := e.Next()
		this.list.Remove(e)
		for ; next != nil; next = next.Next() {
			nextKey := next.Value.(string)
			if val <= this.val[nextKey] {
				newE := this.list.InsertBefore(key, next)
				this.ele[key] = newE
				return
			}
		}
		newE := this.list.PushBack(key)
		this.ele[key] = newE
	}
	return
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if _, ok := this.val[key]; !ok {
		return
	}
	e := this.ele[key]
	if this.val[key] == 1 {
		delete(this.val, key)
		this.list.Remove(e)
		delete(this.ele, key)
	} else {
		this.val[key]--
		val := this.val[key]
		pre := e.Prev()
		this.list.Remove(e)
		for ; pre != nil; pre = pre.Prev() {
			preKey := pre.Value.(string)
			if val >= this.val[preKey] {
				newE := this.list.InsertAfter(key, pre)
				this.ele[key] = newE
				return
			}
		}
		newE := this.list.PushFront(key)
		this.ele[key] = newE
	}
	return
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.list.Len() == 0 {
		return ""
	}
	return this.list.Back().Value.(string)
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.list.Len() == 0 {
		return ""
	}
	return this.list.Front().Value.(string)

}
