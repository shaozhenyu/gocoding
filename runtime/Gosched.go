package main

import (
	"fmt"
	"runtime"
)

func say(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
		runtime.Gosched()
	}
}

func main() {
	go say("111")
	say("222")
}
