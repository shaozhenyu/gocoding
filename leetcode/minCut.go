package main

import "fmt"

func main() {
	s := "aabbaaccccdddaaffa"
	fmt.Println(minCut(s))
}

// aab
func minCut(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	cut := make([]int, len(s) + 1)
	for i := len(s) - 1; i >= 0; i-- {
		cut[i] = len(s)
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j - i <= 2 || dp[i+1][j-1] {
					dp[i][j] = true
					cut[i] = min(cut[i], cut[j+1] + 1)
				}
			}
		}
	}
	return cut[0] - 1
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
