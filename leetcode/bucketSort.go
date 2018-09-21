package main

import "fmt"

func main() {
	bucketSort([]int{10,8,1,4,4,9,7,0,12,3,6})
}

func bucketSort(arr []int) {
	max := -1 << 32
	for i := 0; i < len(arr); i++ {
		max = Max(arr[i], max)
	}
	bucket := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	index := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			arr[index] = i
			index++
			bucket[i]--
		}
	}
	fmt.Println(arr)
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
