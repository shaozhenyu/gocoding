package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

type Test struct {
	Name atomic.Value
	Age  atomic.Value
}

func InitTest() *Test {
	n := new(Test)
	n.Name.Store("szy")
	n.Age.Store(11)
	return n
}

func (t *Test) Store(x1 interface{}, x2 interface{}) {
	t.Name.Store(x1)
	t.Age.Store(x2)
	wg.Done()
}

func (t *Test) Load() {
	name := t.Name.Load().(string)
	age := t.Age.Load().(int)
	fmt.Println(name, age)
	wg.Done()
}

func main() {
	n := InitTest()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go n.Load()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go n.Store("s"+strconv.Itoa(i), i)
	}


	wg.Wait()
}
