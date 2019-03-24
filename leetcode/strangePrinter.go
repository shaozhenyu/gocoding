package main

import "fmt"

func main() {
	s := "baaaacbb"
	fmt.Println(strangePrinter(s))
}

func strangePrinter(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i+1][j] + 1
			}
			for k := i+1; k <= j; k++ {
				if s[k] == s[i] {
					dp[i][j] = min(dp[i][j], dp[i][k-1] + dp[k+1][j])
				}
			}
		}
	}
	return dp[0][len(s)-1]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
