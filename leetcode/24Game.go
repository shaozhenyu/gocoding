package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 8, 7, 4}
	//nums = []int{3, 9, 7, 7}
	//nums = []int{1, 1, 7, 7}
	fmt.Println(judgePoint24(nums))
}

// + - * / ()
// backtrace
func judgePoint24(nums []int) bool {
	arr := make([]float64, 4)
	for i := 0; i < 4; i++ {
		arr[i] = float64(nums[i])
	}
	return backtrace(arr)
}

func backtrace(nums []float64) bool {
	zero := 0.0001
	if len(nums) == 1 {
		if math.Abs(nums[0]-float64(24)) < zero {
			return true
		}
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			p := nums[i]
			q := nums[j]
			add := []float64{p + q, p - q, q - p, p * q}
			if math.Abs(p) > zero {
				add = append(add, q/p)
			}
			if math.Abs(q) > zero {
				add = append(add, p/q)
			}

			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums[:j], nums[j+1:]...)

			for _, v := range add {
				nums = append(nums, v)
				fmt.Println(nums)
				if backtrace(nums) {
					return true
				}
				nums = nums[:len(nums)-1]
			}

			rear := append([]float64{}, nums[j:]...)
			nums = append(nums[:j], q)
			nums = append(nums, rear...)

			rear = append([]float64{}, nums[i:]...)
			nums = append(nums[:i], p)
			nums = append(nums, rear...)
		}
	}
	return false
}
