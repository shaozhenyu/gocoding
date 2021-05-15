package main

import (
	"fmt"
	"container/list"
)

func main() {
	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushFront(5)
	e := l.Back()
	l.Remove(e)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
