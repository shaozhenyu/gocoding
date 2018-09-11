package main

import ()

func foo() *int {
	var x int
	return &x
}

func bar() int {
	x := new(int)
	*x = 1
	return *x
}

func foo1() {
	x1 := make([]int, 8191)
	x2 := make([]int, 8192)
	x3 := make([]int, 8193)
}

func main() {}
