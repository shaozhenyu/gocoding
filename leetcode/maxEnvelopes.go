package main

import (
	"fmt"
	"sort"
)

func main() {
	envelopes := [][]int{[]int{5,4},[]int{6,4},[]int{6,7},[]int{2,3}}
	fmt.Println(maxEnvelopes(envelopes))
}

// binary search
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	
	tail := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		lo, li := 0, len(tail) - 1
		for lo <= li {
			mid := lo + (li - lo)/2
			if envelopes[i][1] > tail[mid] {
				lo = mid + 1
			} else {
				li = mid - 1
			}
		}
		if lo >= len(tail) {
			tail = append(tail, envelopes[i][1])
		} else {
			tail[lo] = envelopes[i][1]
		}
	}
	return len(tail)
}

func maxEnvelopes1(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	max := 1
	dp := make([]int, len(envelopes))
	dp[0] = 1
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				if dp[i] < dp[j] + 1 {
					dp[i] = dp[j] + 1
				}
				if dp[i] > max {
					max = dp[i]
					break
				}
			}
		}
	}
	return max
}

