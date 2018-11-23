package main

import "fmt"

func main() {
	arr := []int{2,1,3,4,4,10, 11, 8}
	fmt.Println(maxChunksToSorted(arr))
}

func maxChunksToSorted(arr []int) int {
	maxLeft := make([]int, len(arr))
	minRight := make([]int, len(arr))
	
	maxLeft[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		maxLeft[i] = max(arr[i], maxLeft[i-1])
	}
	minRight[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		minRight[i] = min(arr[i], minRight[i+1])
	}

	ret := 0
	for i := 0; i < len(arr)-1; i++ {
		if maxLeft[i] <= minRight[i+1] {
			ret++
		}
	}
	return ret + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
