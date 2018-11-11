package main

import (
	"fmt"
	"container/list"
)

func main() {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) []int {
	l := list.New()
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		for e := l.Front(); e != nil; {
			v, _ := e.Value.(int)
			if v < nums[i] {
				tmp := e
				e = e.Next()
				l.Remove(tmp)
			} else {
				e = e.Next()
			}
		}
		l.PushBack(nums[i])
		if i >= k-1 {
			e := l.Front()
			v, _ := e.Value.(int)
			ret = append(ret, v)
			if v == nums[i-k+1] {
				l.Remove(e)
			}
		}
	}
	return ret
}
