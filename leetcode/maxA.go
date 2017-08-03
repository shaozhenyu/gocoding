package main

import "fmt"

func main() {
	fmt.Println(maxA(34))
}

func maxA(N int) int {
	if N <= 6 {
		return N
	}
	dp := make([][]int, N+1)
	// A CA CC CV
	dp[0] = []int{0, 0, 0, 0}
	dp[1] = []int{1, 0, 0, 0}
	dp[2] = []int{2, 1, 0, 0}
	dp[3] = []int{3, 2, 1, 0}

	for i := 4; i <= N; i++ {
		dp[i] = make([]int, 4)
		dp[i][0] = max(dp[i-1][0]+1, dp[i-1][3])
		dp[i][1] = max(dp[i-1][0], dp[i-1][3])
		dp[i][2] = dp[i-1][1]
		dp[i][3] = max(max(dp[i-1][2] * 2, dp[i-2][2] * 3), dp[i-3][2] * 4)
	}

	return dp[N][3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
