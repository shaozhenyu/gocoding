package main

import (
	"fmt"
	"sort"
)

func main() {

}

type TimeMap struct {
	t map[string][]Node
}

type Node struct {
	val string
	timestamp int
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{make(map[string][]Node)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	if _, ok := this.t[key]; !ok {
		this.t[key] = make([]Node, 0)
	}
	this.t[key] = append(this.t[key], Node{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.t[key]; !ok {
		return ""
	}
	v := this.t[key]
	lo, li := 0, len(v)
	for lo <= li {
		mid := lo + (li - lo)/2
		if v[mid].timestamp == timestamp {
			return v[mid].val
		} else if v[mid].timestamp > timestamp {
			li = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if lo < 0 {
		return ""
	}
	if v[lo].timestamp > timestamp {
		lo--
	}
	if lo < 0 {
		return ""
	}
	return v[lo].val
}


