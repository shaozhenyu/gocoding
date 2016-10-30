package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration, n *int) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay * time.Second)
		fmt.Println(text)
		*n++
		close(ch)
	}()
	return ch
}

func main() {
	n := 0
	wait := Publish("publish : haha", 3, &n)
	fmt.Println("I am im Main")
	<-wait
	n++
	fmt.Println("n = ", n)
}
