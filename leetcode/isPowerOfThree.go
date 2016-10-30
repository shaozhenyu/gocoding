package main

import "fmt"

func main() {
	n := 3
	for i := 1; i < 20; i++ {
		n = n * 3
		fmt.Printf("%10d\n", n)
	}
}
