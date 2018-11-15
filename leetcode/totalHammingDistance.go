package main

import "fmt"

func main() {
	nums := []int{4, 14, 2}
	fmt.Println(totalHammingDistance(nums))
}

func totalHammingDistance(nums []int) int {
	total := 0
	for i := 0; i < 32; i++ {
		one := 0
		zero := 0
		for j := 0; j < len(nums); j++ {
			t := nums[j] & 1
			if t == 1 {
				one++
			} else {
				zero++
			}
			nums[j] = nums[j] >> 1
			if nums[j] < 0 {
				nums[j] = 0
			}
		}
		total += one * zero
	}
	return total
}
