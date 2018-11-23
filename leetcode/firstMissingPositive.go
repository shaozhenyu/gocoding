package main

import "fmt"

func main() {
	nums := []int{1,2,0}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
    d := make([]int, len(nums)+2)
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 && nums[i] <= len(d) {
            d[nums[i]] = 1
        }
    }
    for i := 1; i < len(d); i++ {
        if d[i] == 0 {
            return i
        }
    }
    return 1
}
