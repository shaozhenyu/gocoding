package main

import (
	"container/list"
	"fmt"
)

func main() {
	m := Constructor()
	m.AddNum(3)
	m.AddNum(2)
	m.AddNum(1)
	fmt.Println(m.FindMedian())
}

type MedianFinder struct {
	data *list.List
	midl *list.Element
	midr *list.Element
	all map[int]*list.Element
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		data: list.New(),
		midl: nil,
		midr: nil,
		all : make(map[int]*list.Element),
	}
}


func (this *MedianFinder) AddNum(num int) {
	var ne *list.Element
	if e, ok := this.all[num]; ok {
		ne = this.data.InsertBefore(num, e)
	} else {
		e := this.data.Front()
		for ; e != nil; e = e.Next() {
			v := e.Value.(int)
			if num <= v {
				ne = this.data.InsertBefore(num, e)
				break
			}
		}
		if e == nil {
			ne = this.data.PushBack(num)
		}
	}
	this.all[num] = ne

	if this.midl == nil && this.midr == nil {
		this.midl = this.data.Front()
	} else {
		tr := this.midr
		if num <= this.midl.Value.(int) {
			t := this.midl
			this.midl = t.Prev()
			this.midr = t
		} else {
			this.midr = this.midl.Next()
		}
		if tr != nil {
			this.midl = this.midr
			this.midr = nil
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.midl == nil {
		return 0
	}
	if this.midr == nil {
		return float64(this.midl.Value.(int))
	}
	return (float64(this.midl.Value.(int)) + float64(this.midr.Value.(int)))/2.0
}


