package main

import "fmt"

func main() {
	nums := []int{1,2,3,3,4,4,5,5}
	fmt.Println(isPossible(nums))
}

func isPossible(nums []int) bool {
	have := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		have[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		if have[nums[i]] > 0 {
			start := nums[i]
			count := 0
			for have[start] > 0 && have[start] > have[start-1] {
				have[start]--
				start++
				count++
			}
			if count < 3 {
				return false
			}
		}
	}
	return true
}
