package main

import "fmt"

func main() {
	nums := []int{3, 2, 3, 3}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	fmt.Println("nums: ", nums)
	ones := 0
	twos := 0
	mask := 0
	for i := 0; i < len(nums); i++ {
		twos = twos | (ones & nums[i])
		ones = ones ^ nums[i]
		fmt.Println(ones & twos)
		mask = ^(ones & twos)
		fmt.Println("one : ", nums[i], ones, twos, mask)

		ones = ones & mask
		twos = twos & mask
		fmt.Println("two : ", nums[i], ones, twos, mask)
	}
	return ones
}
