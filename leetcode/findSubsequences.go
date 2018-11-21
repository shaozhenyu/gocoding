package main

import "fmt"

func main() {
	nums := []int{4, 6, 7, 7} //, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(findSubsequences(nums))
}

func findSubsequences(nums []int) [][]int {
	n := make([]int, len(nums))
	return findSubsequence(0, 0, n, nums, [][]int{})
}

func findSubsequence(p int, size int, n []int, nums []int, ret [][]int) [][]int {
	if size >= 2 {
		newN := make([]int, size)
		copy(newN, n[:size])
		ret = append(ret, newN)
	}

	for i := p; i < len(nums); i++ {
		if i-1 >= p && nums[i-1] == nums[i] {
			continue
		}
		if size == 0 || n[size-1] <= nums[i] {
			n[size] = nums[i]
			ret = findSubsequence(i+1, size+1, n, nums, ret)
			n[size] = 0
		}
	}
	return ret
}

func findSubsequences1(nums []int) [][]int {
	ret := [][]int{}
	for i := 0; i < len(nums); i++ {
		all := [][]int{}
		m := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				continue
			}

			old := [][]int{}
			if _, ok := m[nums[j]]; !ok {
				for k := 0; k < len(all); k++ {
					if nums[j] >= all[k][len(all[k])-1] {
						old = append(old, all[k])
						all[k] = append(all[k], nums[j])
					}
				}
			} else {
				for k := 0; k < len(all); k++ {
					if nums[j] == all[k][len(all[k])-1] {
						old = append(old, all[k])
						all[k] = append(all[k], nums[j])
					}
				}
			}

			if _, ok := m[nums[j]]; !ok {
				all = append(all, []int{nums[i], nums[j]})
			}
			all = append(all, old...)
			m[nums[j]] = 1
		}
		ret = append(ret, all...)
	}

	return ret
}
