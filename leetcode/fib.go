package main

import "fmt"

func main() {
	fmt.Println(fib(10))
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-2) + fib(N-1)
}
