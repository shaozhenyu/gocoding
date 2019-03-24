package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,3,2,2,3,1}
	nums = []int{1,1,2,1,2,2,1}
	//nums = []int{4,5,5,6}
	nums = []int{5,3,1,2,6,7,8,5,5}
	wiggleSort(nums)
	fmt.Println(nums)
}

func wiggleSort(nums []int) {
	sort.Ints(nums)
	if len(nums) <= 2 {
		return
	}
	mid := len(nums)/2
	l := len(nums) - 1
	c := len(nums) - 1
	if len(nums)%2 == 1 {
		first := nums[0]
		for i := 1; i < len(nums); i++ {
			nums[i-1] = nums[i]
		}
		nums[len(nums)-1] = first
		l--
		c--
	}
	from, to := 0, 0
	tmp := nums[from]
	for l >= 0 {
		if from < mid {
			to = c - 1 - 2 * from
		} else {
			to = 2 * (c - from) + 1
		}
		fmt.Println(from, to)
		tmp1 := nums[to]
		nums[to] = tmp
		tmp = tmp1
		from = to
		l--
	}
}

func wiggleSort1(nums []int)  {
	sort.Ints(nums)
	if len(nums) <= 2 {
		return
	}
	t := make([]int, len(nums))
	copy(t, nums)
	idx := len(nums)-1
	mid := len(nums)/2
	i, j := 0, mid
	if len(nums)%2 == 1 {
		nums[idx] = t[i]
		i++
		j++
		idx--
	}
	for idx >= 1 {
		nums[idx] = t[j]
		idx--
		j++
		nums[idx] = t[i]
		idx--
		i++
	}
}

