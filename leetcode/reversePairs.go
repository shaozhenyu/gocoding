package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{100,200, 2, 20, 15, 1, 7, -1, 3, 8}
	fmt.Println(reversePairs(nums))
}

func reversePairs(nums []int) int {
	fmt.Println(nums)
	coy := make([]int, len(nums))
	copy(coy, nums)
	sort.Ints(coy)
	fmt.Println(coy)
	bit := make([]int, len(nums)+1)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] * 2 + 1 <= coy[len(coy)-1] {
			count += search(index(nums[i]*2+1, coy), bit)
		}
		insert(index(nums[i], coy), bit)
		fmt.Println(bit)
	}
	return count
}

func search(idx int, bit []int) int {
	ret := 0
	for idx < len(bit) {
		fmt.Println("search:", idx)
		ret += bit[idx]
		idx += (idx & (-1 * idx))
	}
	return ret
}

func insert(idx int, bit []int) {
	for idx >= 1 {
	fmt.Println("insert: ", idx)
		bit[idx]++
		idx -= (idx & (-1 * idx))
	}
}

func index(val int, nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] > val {
			right = mid - 1
		} else if nums[mid] < val {
			left = mid + 1
		} else {
			return mid + 1
		}
	}
	return left + 1
}
