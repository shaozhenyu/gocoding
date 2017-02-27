package main

import (
	"fmt"
	"container/list"
)

func main() {
	ret := nextGreaterElements([]int{100, 99, 102, 98})
	fmt.Println(ret)
}

func nextGreaterElements(nums []int) []int {
	lenN := len(nums)
	ret := make([]int, lenN)
	for i := 0; i < lenN; i++ {
		ret[i] = -1
	}
	l := list.New()

	for i := 0; i < lenN * 2; i++ {
		num := nums[i%lenN]
		for l.Len() != 0 {
			v, _ := l.Back().Value.(int)
			if nums[v] < num {
				ret[v] = num
				top := l.Back()
				l.Remove(top)
			} else {
				break
			}
		}
		if i < lenN {
			l.PushBack(i)
		}
	}
	return ret
}
