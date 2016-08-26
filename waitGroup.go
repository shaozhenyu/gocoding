package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	num = 10000000
)

func main() {
	TestFunc("testchan", TestChan)
}

func TestFunc(name string, f func()) {
	st := time.Now().UnixNano()
	f()
	fmt.Printf("task %s cost %d \r\n", name, (time.Now().UnixNano()-st)/int64(time.Millisecond))
}

func TestChan() {
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(1)

	go func() {
		for _ = range c {
		}
		wg.Done()
	}()

	for i := 0; i < num; i++ {
		c <- "123"
	}

	close(c)
	wg.Wait()
}
