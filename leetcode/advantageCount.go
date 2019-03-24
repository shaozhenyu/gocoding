package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{8,2,4,4,5,6,6,0,4,7}
	B := []int{0,8,7,4,4,2,8,5,2,0}
	fmt.Println(advantageCount(A, B))
}

func advantageCount(A []int, B []int) []int {
	sort.Ints(A)
	ret := make([]int, len(A))
	for i := 0; i < len(B); i++ {
		idx := binarySearch(B[i], A)
		for idx < len(A) && A[idx] <= B[i] {
			idx++
		}
		if idx == len(A) {
			ret[i] = A[0]
			A = A[1:]
		} else {
			ret[i] = A[idx]
			A = append(A[:idx], A[idx+1:]...)
		}
	}
	return ret
}

func binarySearch(v int, nums []int) int {
	lo, li := 0, len(nums) - 1
	for lo <= li {
		mid := lo + (li - lo)/2
		if v == nums[mid] {
			return mid
		} else if v > nums[mid] {
			lo = mid + 1
		} else {
			li = mid - 1
		}
	}
	return lo
}
