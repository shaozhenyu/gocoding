package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	val  map[int]int
	rv   []int
	size int
}

//["RandomizedSet","remove","remove","insert","getRandom","remove","insert"]
//[[],[0],[0],[0],[],[0],[0]]

//["RandomizedSet","insert","insert","getRandom","getRandom","insert","remove","getRandom","getRandom","insert","remove"]
//[[],[3],[3],[],[],[1],[3],[],[],[0],[0]]

func main() {
	r := Constructor()
	r.Insert(3)
	r.Insert(3)
	fmt.Println(r.GetRandom())
	fmt.Println(r.GetRandom())
	r.Insert(1)
	r.Remove(3)
	fmt.Println(r.GetRandom())
	fmt.Println(r.GetRandom())
	r.Insert(0)
	r.Remove(0)
	fmt.Println(r)
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	println("111")
	return RandomizedSet{
		val: make(map[int]int),
		//pos: make(map[int]int),
		rv:   []int{},
		size: 0,
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	println("insert")
	if _, ok := this.val[val]; ok {
		return false
	}
	this.val[val] = this.size
	this.rv = append(this.rv, val)
	//this.pos[this.size] = val
	this.size = this.size + 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	println("remove")
	if pos, ok := this.val[val]; ok {
		fmt.Println("11")
		fmt.Println(this.rv)
		rv := this.rv
		if len(rv) == 1 {
			rv = []int{}
		} else {
			rv[pos] = rv[len(rv)-1]
			this.val[rv[pos]] = pos
			rv = rv[:len(rv)-1]
		}
		delete(this.val, val)
		this.size = this.size - 1
		this.rv = rv
		fmt.Println("33")
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	println("random")
	if len(this.val) == 0 {
		return 0
	}
	//fmt.Println(this.val, this.pos, this.size)
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(this.size)
	return this.rv[x]
}
