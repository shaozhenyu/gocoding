package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberToGenerate = 1000000
)

func main() {
	start := time.Now()
	fmt.Println("Start")

	list := lottoNumbers(numberToGenerate)

	d := time.Now().Sub(start)
	fmt.Println("End ", len(list), " ", d.Seconds())
}

func lottoNumbers(n int) [][]int {
	var list = [][]int{}
	rand.Seed(time.Now().UnixNano())

	rec := make(chan []int, n)

	go func(c chan []int) {
		for {
			list = append(list, <-c)
		}
	}(rec)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				add(rec)
			}
		}()
	}
	wg.Wait()

	return list
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
