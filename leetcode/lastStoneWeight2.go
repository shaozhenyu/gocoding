package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastStoneWeightII([]int{4,3,3,1,1}))
}

func lastStoneWeightII(stones []int) int {
	cur := map[int]struct{}{
		stones[0]: struct{}{},
		-1 * stones[0]: struct{}{},
	}
	for i := 1; i < len(stones); i++ {
		t := make(map[int]struct{})
		for k := range cur {
			t[k + stones[i]] = struct{}{}
			t[k - stones[i]] = struct{}{}
		}
		cur = t
	}
	fmt.Println(cur)
	ret := 2 << 31
	for k := range cur {
		ret = min(ret, abs(k))
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
