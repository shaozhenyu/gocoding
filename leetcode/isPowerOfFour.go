package main

import "fmt"

func main() {
	n := 4
	for i := 1; i < 20; i++ {
		n = n * 4
		fmt.Printf("%10d\n", n)
	}
}
