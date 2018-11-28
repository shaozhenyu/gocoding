package main

import "fmt"

func main() {
	S := "abbaabbaa"
	fmt.Println(countPalindromicSubsequences(S))
}

// abba a b aa bb aba abba
func countPalindromicSubsequences(S string) int {
	dp := make([][]int, len(S))
	mod := 1000000000 + 7
	for i := 0; i < len(S); i++ {
		dp[i] = make([]int, len(S))
		dp[i][i] = 1
	}
	for step := 1; step < len(S); step++ {
		for i := 0; i + step < len(S); i++ {
			j := i + step
			if S[i] == S[j] {
				left := i + 1
				right := j - 1
				for left <= right && S[i] != S[left] {
					left++
				}
				for left <= right && S[j] != S[right] {
					right--
				}
				if left > right {
					dp[i][j] = (dp[i+1][j-1] * 2 + 2)%mod
				} else if left == right {
					dp[i][j] = (dp[i+1][j-1] * 2 + 1)%mod
				} else {
					dp[i][j] = (dp[i+1][j-1] * 2 - dp[left+1][right-1])%mod
				}
			} else {
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1])%mod
			}
		}
	}
	return dp[0][len(S)-1]
}
