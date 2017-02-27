package main

import (
	"fmt"
	"sort"
)

func main() {
	findRelativeRanks([]int{1,2,3,5,4,3,2,1})
}

func findRelativeRanks(nums []int) []string {
	lenN := len(nums)
	//m := map[int]int{}
	m := make([]int, lenN)
	for i := 0; i < lenN; i++ {
		m[i] = nums[i]
	}
	ret := make([]int, lenN)
	lenM := lenN
	sort.Ints(nums)
	for i := lenN-1; i >= 0; i-- {
		for j := 0; j < lenM; j++ {
			if nums[i] == m[j] {
				ret[j] = lenN - i
				m[j] = -99999
				break
			}
		}
		//for k, v := range(m) {
		//	if nums[i] == v {
		//		ret[k] = lenN - i
		//		delete(m, k)
		//		break
		//	}
		//}
	}
	fmt.Println(ret)
	return []string{}
}
