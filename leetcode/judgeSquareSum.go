package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(5))
}

func judgeSquareSum(c int) bool {
	for i := 0; i <= c/2; i++ {
		if i * i > c {
			break
		}
		t := c - i * i
		k := int(math.Sqrt(float64(t)))
		if k * k == t {
			return true
		}
	}
	return false
}
