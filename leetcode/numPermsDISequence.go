package main

import "fmt"

func main() {
	s := "DIDDI"
	fmt.Println(numPermsDISequence(s))
}

func numPermsDISequence(S string) int {
	mod := 1000000007
	dp := make([][]int, len(S)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(dp))
	}
	for i := 0; i < len(dp); i++ {
		dp[0][i] = 1
	}
	for i := 0; i < len(S); i++ {
		t := 0
		if S[i] == 'D' {
			for j := len(S) - i - 1; j >= 0; j-- {
				t = (t + dp[i][j+1]) % mod
				dp[i+1][j] = t
			}
		} else {
			for j := 0; j < len(S) - i; j++ {
				t = (t + dp[i][j]) % mod
				dp[i+1][j] = t
			}
		}
	}
	return dp[len(S)][0]
}
