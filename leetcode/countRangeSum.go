package main

import "fmt"

func main() {

}

func countRangeSum(nums []int, lower int, upper int) int {
	ret := 0
	for i := 0; i < len(nums); i += 2 {
		v := nums[i]
		if v >= lower && v <= upper {
			ret++
		}
		for j := i + 1; j < len(nums); j++ {
			v += nums[j]
			if v >= lower && v <= upper {
				ret++
			}
			t := v - nums[i]
			if t >= lower && t <= upper {
				ret++
			}
		}
	}
	return ret
}
