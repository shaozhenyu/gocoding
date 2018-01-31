package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var list = [][]int{}

func main() {

	rand.Seed(time.Now().UnixNano())

	rec := make(chan []int, 10000000)

	go func(c chan []int) {
		for {
			list = append(list, <-c)
		}
	}(rec)

	start := time.Now()
	fmt.Println("Start")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				add(rec)
			}
		}()
	}
	wg.Wait()
	d := time.Now().Sub(start)
	fmt.Println("End ", len(list), " ", d.Seconds())
}

func add(c chan []int) {
	c <- []int{
		rand.Intn(49),
		rand.Intn(49),
		rand.Intn(49),
		rand.Intn(49),
		rand.Intn(49),
		rand.Intn(49),
		rand.Intn(49),
	}
}
