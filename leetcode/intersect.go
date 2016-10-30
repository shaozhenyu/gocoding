package main

import (
	"fmt"
	"sort"
)

func main() {

	n1 := []int{1, 2, 2, 1, 3}
	n2 := []int{3, 2, 2}
	n := intersect(n1, n2)
	fmt.Println(n)
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	nums := []int{}

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			nums = append(nums, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}

	return nums
}
