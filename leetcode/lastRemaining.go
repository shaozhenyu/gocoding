package main

import "fmt"

func main() {
	for i := 0; i <= 20; i += 1 {
		fmt.Println(i, lastRemaining(i))
	}
}

func lastRemaining(n int) int {
	head := 1
	step := 1
	left := true
	for n > 1 {
		if left || n%2 == 1 {
			head += step
		}
		n = n/2
		step *= 2
		left = !left
	}
	return head
}

func lastRemaining1(n int) int {
	if n <= 2 {
		return n
	}
	ret := make([]int, n/2)
	num := 2
	for i := 0; i < n/2; i++ {
		ret[i] = num
		num += 2
	}

	leftFlag := false
	for len(ret) > 1 {
		//fmt.Println(ret)
		if leftFlag {
			ret = left(ret)
			leftFlag = false
		} else {
			ret = right(ret)
			leftFlag = true
		}
	}
	return ret[0]
}

func left(nums []int) []int {
	//fmt.Println("left", nums)
	ret := make([]int, len(nums)/2)
	pos := 0
	for i := 1; i < len(nums); i += 2 {
		ret[pos] = nums[i]
		pos++
	}
	return ret
}

func right(nums []int) []int {
	//fmt.Println("right:",nums)
	ret := make([]int, len(nums)/2)
	pos := len(ret) - 1
	for i := len(nums)-2; i >= 0; i -= 2 {
		ret[pos] = nums[i]
		pos--
	}
	return ret
}
