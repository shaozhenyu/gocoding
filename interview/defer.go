package main

import (
	"fmt"
)

func main() {
	defer_call()
}

//defer执行顺序是先进后出
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
