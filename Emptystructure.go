package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

var itemExists = struct{}{}

type HSet struct {
	items map[interface{}]struct{}
}

func main() {
	hSet := New()
}

func New() *HSet {
	return &HSet{items: make(map[interface{}]struct{})}
}

func (hSet *HSet) Add(item interface{}) {
	hSet.items[item] = itemExists
}

func (hSet *HSet) Remove(item interface{}) {
	delete(hSet.items, item)
}

func (hSet *HSet) Contains(item interface{}) bool {
	if _, ok := hSet.items[item]; ok {
		return true
	}
	return false
}
