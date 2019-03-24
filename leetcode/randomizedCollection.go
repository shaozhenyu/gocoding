package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := Constructor()
	c.Insert(1)
	c.Remove(1)
	c.Insert(2)
	fmt.Println(c.vals)
	c.Remove(1)
	fmt.Println(c.vals)
	fmt.Println(c.GetRandom())
}

type RandomizedCollection struct {
	rd *rand.Rand
	m map[int]map[int]bool
	vals []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		rd: rand.New(rand.NewSource(time.Now().Unix())),
		m: make(map[int]map[int]bool),
		vals : make([]int, 0, 4096),
	}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	exists := false
	if _, ok := this.m[val]; ok {
		exists = true
	} else {
		this.m[val] = make(map[int]bool)
	}
	this.vals = append(this.vals, val)
	this.m[val][len(this.vals)-1] = true
	return exists
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	for idx := range this.m[val] {
		delete(this.m[val], idx)
		last := this.vals[len(this.vals)-1]
		if idx != len(this.vals)-1 {
			delete(this.m[last], len(this.vals)-1)
			this.vals[idx] = last
			this.m[last][idx] = true
		}
		this.vals = this.vals[:len(this.vals)-1]
		return true
	}
	return false
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	if len(this.vals) == 0 {
		return 0
	}
	idx := this.rd.Int()%len(this.vals)
	return this.vals[idx]
}

