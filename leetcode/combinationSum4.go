package main

import (
	"fmt"
	"sort"
)

//28730

func main() {

	nums := []int{1, 50}
	target := 200
	for ; target <= 200; target++ {
		fmt.Println(combinationSum4(nums, target))
	}
}

func combinationSum4(nums []int, target int) int {
	sum := f(nums, target, []int{})
	return sum
}

func f(nums []int, target int, all []int) int {
	r := 0
	if target < 0 {
		return 0
	}
	if target == 0 {
		fmt.Println("all:", all)
		sum := 0
		ret := 1
		for i := 0; i < len(all); i++ {
			sum += all[i]
		}
		sort.Ints(all)
		for i := 0; i < len(all)-1; i++ {
			ret *= get(sum, all[i])
			fmt.Println("ret : ", ret)
			sum -= all[i]
		}
		return ret
	}
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= target/nums[i]; j++ {
			newAll := make([]int, len(all)+1)
			copy(newAll, all)
			newAll[len(all)] = j
			r += f(nums[i+1:], target-j*nums[i], newAll)
		}
	}
	return r
}

func get(sum int, count int) int {
	fmt.Println(sum, count)
	if sum == count {
		return 1
	}
	if (sum - count) == 1 {
		return sum
	}
	ret := 1
	for i := 0; i < count; i++ {
		ret = ret * (sum - i) / (1 + i)
		fmt.Println(ret)
	}
	return ret
}
