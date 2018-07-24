package main

import "fmt"

func main() {
	nums := []int{11, 11, 11, 12}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	sum := 1
	for i := 0; i < len(nums); i++ {
		sum = sum ^ nums[i]
	}
	return sum
}

func singleNumber1(nums []int) []int {
	length := len(nums)
	p := map[int]int{}

	for i := 0; i < length; i++ {
		if _, ok := p[nums[i]]; ok {
			delete(p, nums[i])
		} else {
			p[nums[i]] = 1
		}
	}
	
	i := 0
	array := make([]int, 2)
	for k, _ := range p {
		array[i] = k;
		i++
	}
	return array
}
