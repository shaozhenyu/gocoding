package main

import "fmt"

func main() {
	weights := []int{1,2,3,1,1}
	D := 4
	fmt.Println(shipWithinDays(weights, D))
}

func shipWithinDays(weights []int, D int) int {
	low, high := -2 << 31, 0
	for i := 0; i < len(weights); i++ {
		high += weights[i]
		if weights[i] > low {
			low = weights[i]
		}
	}
	for low < high {
		mid := low + (high - low)/2
		weight := 0
		day := 0
		for i := 0; i < len(weights); i++ {
			if weights[i] + weight > mid {
				day++
				weight = 0
			}
			weight += weights[i]
		}
		day++
		if day <= D {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
