package main

import "fmt"

func main() {
	fmt.Println(climbStairs(44))
}

func climbStairs(n int) int {

	if n < 0 {
		return 0
	}

	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	a := 1
	b := 2
	c := 0
	for n > 2 {
		c = a + b
		a, b = b, c
		n--
	}
	return c
}
