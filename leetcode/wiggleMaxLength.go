package main

import "fmt"

func main() {
	nums := []int{1,17,5,10,13,15,10,5,16,8}
	fmt.Println(wiggleMaxLength(nums))
}

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 2 {
		return 2
	}
	flag := 0
	count := 1
	for i := len(nums) - 2; i >= 0; i-- {
		de := nums[i+1] - nums[i]
		if de > 0 && flag == -1 {
			count++
			flag = 1
			continue
		}
		if de < 0 && flag == 1 {
			count++
			flag = -1
			continue
		}
		if flag == 0 && de != 0{
			if de < 0 {
				flag = -1
			} else {
				flag = 1
			}
			count++
		}
	}
	return count
}
