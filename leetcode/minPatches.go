package main

import "fmt"

func main() {
	nums := []int{1,3,7}
	n := 123125
	fmt.Println(minPatches(nums, n))
}

func minPatches(nums []int, n int) int {
	
}

func minPatches(nums []int, n int) int {
	all := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		t := make(map[int]bool)
		t[nums[i]] = true
		for j := i + 1; j < len(nums); j++ {
			t1 := make(map[int]bool)
			for k := range t {
				t1[k] = true
				t1[k+nums[j]] = true
			}
			t = t1
		}
		for k := range t {
			all[k] = true
		}
	}
	return 0
}
