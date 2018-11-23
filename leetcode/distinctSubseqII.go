package main

import "fmt"

func main() {
	S := "aacbaaba"
	fmt.Println(distinctSubseqII(S))
}

func distinctSubseqII(S string) int {
	dp := make([]int, len(S))
	for i := len(S) - 1; i >= 0; i-- {
		for j := i + 1; j < len(S); j++ {
			if S[i] != S[j] {
				dp[i] += dp[j]
			}
		}
		dp[i] += 1
	}
	sum := 0
	fmt.Println(dp)
	for i := 0; i < len(dp); i++ {
		sum += dp[i]
	}
	return sum
}
