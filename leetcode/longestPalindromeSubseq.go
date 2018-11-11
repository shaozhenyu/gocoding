package main

import "fmt"

func main() {
	s := "abcccceffdccccba"
	fmt.Println(longestPalindromeSubseq(s))
}

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
