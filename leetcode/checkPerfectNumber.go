package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPerfectNumber(6))
}

func checkPerfectNumber(num int) bool {
	sum := 1
	sq := int(math.Sqrt(float64(num))) + 1
	for i := 2; i < sq; i++ {
		if num%i != 0 {
			continue
		}
		j := num/i
		fmt.Println(i, j)
		sum += i
		if j != i {
			sum += j
		}
	}
	return sum == num
}
