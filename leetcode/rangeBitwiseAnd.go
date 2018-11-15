package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(4, 7))
}

func rangeBitwiseAnd(m int, n int) int {
	var pos uint
	for m != n {
		m >>= 1
		n >>= 1
		pos++
	}
	return 1 << pos
}
