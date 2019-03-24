package main

import (
	"fmt"
	"container/list"
)

func main() {
	A := []int{1,2,1,2,3}
	K := 2
	fmt.Println(subarraysWithKDistinct(A, K), 7)
	A = []int{1,2,2,1,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,2,3,1,4,3,1,2,5,1,2,3,4,1,2,3,4,5,3,3,2,1,2,3,4,3,1,2,3}
	// A = []int{1,2,3,4,1,2,2,3,2,1,2,3,4,2,2,1,2,3,3,2,1,2,3,4,2,1,2}
	// A = []int{1,2,1,3,4}
	K = 3

	A = []int{2,2,1,2,2,2,1,1}
	K = 2
	fmt.Println(subarraysWithKDistinct(A, K), 23)
	A = []int{5,7,5,2,3,3,4,1,5,2,7,4,6,6,3,3,4,4,7}
	K = 7
	fmt.Println(subarraysWithKDistinct(A, K), 52)
}

func subarraysWithKDistinct(A []int, K int) int {
	fmt.Println(A)
	mlist := make(map[int]*list.List)
	for i := 0; i < len(A); i++ {
		if _, ok := mlist[A[i]]; !ok {
			mlist[A[i]] = list.New()
		}
		mlist[A[i]].PushBack(i)
	}
	sum := 0
	startIdx, lastIdx := -1, -1
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		j := i
		if i > 0 {
			m[A[i-1]] -= 1
			if m[A[i-1]] > 0 {
				e := mlist[A[i-1]].Front()
				idx := e.Value.(int)
				if idx > startIdx {
					startIdx = idx
				}
			} else if m[A[i-1]] == 0 {
				delete(m, A[i-1])
				startIdx = -1
			}
		}
		if lastIdx != -1 {
			j = lastIdx
		}
		for ; j < len(A); j++ {
			m[A[j]]++
			if len(m) == K {
				if startIdx == -1 {
					startIdx = j
				}
				lastIdx = j + 1
			} else if len(m) > K {
				delete(m, A[j])
				break
			}
		}
		if lastIdx == -1 || len(m) < K {
			return sum
		}
		sum += lastIdx - startIdx
		mlist[A[i]].Remove(mlist[A[i]].Front())
	}
	return sum
}
