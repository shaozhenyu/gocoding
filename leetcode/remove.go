package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 3, 4}

	//len := removeElement(nums, 3)
	//fmt.Println(len)

	len := removeDuplicates(nums)
	fmt.Println(len)
}

func removeDuplicates(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	before := nums[0]

	for i := 1; i < length; i++ {
		if nums[i] == before {
			length--
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			before = nums[i]
		}
	}
	fmt.Println(nums)
	return length
}

//func removeDuplicates(nums []int) int {
//	length := len(nums)
//
//	if length == 0 {
//		return 0
//	}
//
//	before := nums[0]
//
//	for i := 1; i < length; i++ {
//		if nums[i] == before {
//			length--
//			m := i
//			for j := i+1; j < len(nums); j++ {
//				nums[m] = nums[j]
//				m++
//			}
//
//			i--
//		} else {
//			before = nums[i]
//		}
//	}
//
//	fmt.Println(nums)
//
//	return length
//}

func removeElement(nums []int, val int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		fmt.Println(length, nums[i], i)
		if val == nums[i] {
			length--
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return length
}
