package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(grayCode(4))
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	num := int(math.Pow(float64(2), float64(n)))
	ret := make([]int, num)
	ret[0] = 0
	ret[1] = 1
	cyc := 2
	for i := 2; i < num; {
		for j := cyc - 1; j >= 0; j-- {
			fmt.Println(i, j)
			ret[i] = cyc + ret[j]
			i++
		}
		cyc *= 2
	}
	return ret
}
