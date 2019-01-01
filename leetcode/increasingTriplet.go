package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 0,0,0,0,0,10000000}
	fmt.Println(increasingTriplet(nums))
}

func increasingTriplet(nums []int) bool {
	min1 := 2 << 31
	min2 := 2 << 31
	for _, v := range nums {
		if v <= min1 {
			min1 = v
		} else if v <= min2 {
			min2 = v
		} else {
			return true
		}
	}
	return false
}
