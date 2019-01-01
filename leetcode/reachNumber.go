package main

import (
	"fmt"
	"math"
)

func main() {
	target := 10
	fmt.Println(reachNumber(target))
}

//  (1 + n) * n / 2 = target
func reachNumber(target int) int {
	if target < 0 {
		target = -1 * target
	}
	t := int(math.Sqrt(float64(target * 2)))
	for t*(t+1)/2 >= target {
		t--
	}
	t++
	v := (1 + t) * t / 2
	de := v - target
	for de%2 != 0 {
		t++
		de += t
	}
	return t
}
