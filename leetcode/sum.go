package main

import "fmt"

func main() {
	fmt.Println(Sum(214748647, 1))
}

func Sum(a, b int) int {
	c := a + b
	//-2147483648－－2147483647
	fmt.Println(c)
	return a + b
}
