package main

import "fmt"

func main() {
	nums := []int{2,3,1,1,4}
	nums = []int{8,2,4,4,4,9,5,2,5,8,8,0,8,6,9,1,1,6,3,5,1,2,6,6,0,4,8,6,0,3,2,8,7,6,5,1,7,0,3,4,8,3,5,9,0,4,0,1,0,5,9,2,0,7,0,2,1,0,8,2,5,1,2,3,9,7,4,7,0,0,1,8,5,6,7,5,1,9,9,3,5,0,7,5}
	// nums = []int{2,3,1,1,4,2,3,4,1,1,2,3,4,5,3,1,2,3,1,6,1,1,1,1,5}
	fmt.Println(jump(nums))
	fmt.Println(jump1(nums))
}

type Node struct {
	idx int
	step int
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	start, end := 1, nums[0]
	endStep := make([]int, 0)
	endStep = append(endStep, end)
	for start <= end {
		if end >= len(nums) - 1 {
			return len(endStep)
		}
		v := start + nums[start]
		if v > end {
			for i := 0; i < len(endStep); i++ {
				if start <= endStep[i] {
					endStep = endStep[:i+1]
				}
			}
			endStep = append(endStep, v)
			end = v
		}
		start++
	}
	return -1
}

func jump1(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	queue := make([]Node, 1, len(nums))
	exists := make(map[int]int)
	start, end := 0, 0
	queue[end] = Node{0, 0}
	exists[start] = 0
	end++
	for start <= end {
		now := queue[start]
		start++
		for i := 1; i <= nums[now.idx]; i++ {
			t := now.idx + i
			if t >= len(nums) - 1 {
				return now.step + 1
			}
			if _, ok := exists[t]; ok {
				continue
			}
			queue = append(queue, Node{t, now.step+1})
			exists[t] = now.step + 1
			end++
		}
	}
	return -1
}
