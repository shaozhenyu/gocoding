package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2, 0, 0}
	nums = []int{1,-2,-5,-4,-3,3,3,5}
	target := -11
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		sum := nums[i]
		if sum > target && nums[i] > 0 {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}
			if sum + nums[j] > target && nums[j] > 0 {
				break
			}
			sum += nums[j]
			for x := j + 1; x < len(nums); x++ {
				if x > j + 1 && nums[x] == nums[x-1] {
					continue
				}
				if sum + nums[x] > target && nums[x] > 0 {
					break
				}
				sum += nums[x]
				for y := x + 1; y < len(nums); y++ {
					if y > x + 1 && nums[y] == nums[y-1] {
						continue
					}
					if sum + nums[y] > target && nums[y] > 0 {
						break
					}
					if sum + nums[y] == target {
						ret = append(ret, []int{nums[i], nums[j], nums[x], nums[y]})
					}
				}
				sum -= nums[x]
			}
			sum -= nums[j]
		}
	}
	return ret
}
