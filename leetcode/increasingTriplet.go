package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 5, 3, 2}
	fmt.Println(increasingTriplet(nums))
}

func increasingTriplet(nums []int) bool {
	min1 := 999999
	min2 := 999999
	for _, v := range nums {
		if v < min1 {
			min1 = v
		} else if v < min2 {
			min2 = v
		} else {
			return true
		}
	}
	return false
}
