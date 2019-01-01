package main

import "fmt"

func main() {
	nums := []int{1, 6, 3, 6}
	nums = []int{5, 5, 5, 5}
	fmt.Println(judgePoint24(nums))
}

func judgePoint24(nums []int) bool {
	if check(calTwo(nums[0], 1, nums[1], 1), calTwo(nums[2], 1, nums[3], 1)) {
		return true
	}
	if check(calTwo(nums[0], 1, nums[2], 1), calTwo(nums[1], 1, nums[3], 1)) {
		return true
	}
	if check(calTwo(nums[0], 1, nums[3], 1), calTwo(nums[1], 1, nums[2], 1)) {
		return true
	}
	if check([][]int{[]int{nums[0], 1}}, calThree(nums[1], 1, nums[2], 1, nums[3], 1)) {
		return true
	}
	if check([][]int{[]int{nums[1], 1}}, calThree(nums[0], 1, nums[2], 1, nums[3], 1)) {
		return true
	}
	if check([][]int{[]int{nums[2], 1}}, calThree(nums[1], 1, nums[0], 1, nums[3], 1)) {
		return true
	}
	if check([][]int{[]int{nums[3], 1}}, calThree(nums[1], 1, nums[2], 1, nums[0], 1)) {
		return true
	}
	return false
}

func calThree(x1, x2, y1, y2, z1, z2 int) [][]int {
	ret := [][]int{}
	for _, t1 := range calTwo(x1, x2, y1, y2) {
		ret = append(ret, calTwo(t1[0], t1[1], z1, z2)...)
	}
	for _, t1 := range calTwo(x1, x2, z1, z2) {
		ret = append(ret, calTwo(t1[0], t1[1], y1, y2)...)
	}
	for _, t1 := range calTwo(y1, y2, z1, z2) {
		ret = append(ret, calTwo(t1[0], t1[1], x1, x2)...)
	}
	return ret
}

func calTwo(x1, x2, y1, y2 int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, []int{x1 * y2 + y1 * x2, x2 * y2})
	ret = append(ret, []int{x1 * y1, x2 * y2})
	ret = append(ret, []int{x1 * y2 - y1 * x2, x2 * y2})
	ret = append(ret, []int{y1 * x2 - x1 * y2, x2 * y2})
	if y1 != 0 {
		ret = append(ret, []int{x1 * y2, x2 * y1})
	}
	if x1 != 0 {
		ret = append(ret, []int{x2 * y1, x1 * y2})
	}
	return ret
}

func check(x1 [][]int, y1 [][]int) bool {
	for _, x := range x1 {
		for _, y := range y1 {
			if checkOne(calTwo(x[0], x[1], y[0], y[1])) {
				return true
			}
		}
	}
	return false
}

func checkOne(x [][]int) bool {
	for _, v := range x {
		if v[0]%v[1] == 0 && v[0]/v[1] == 24 {
			return true
		}
	}
	return false
}
