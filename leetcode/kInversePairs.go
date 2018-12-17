package main

import "fmt"

func main() {
	n := 5
	k := 7
	fmt.Println(kInversePairs(n, k))
}
// i < j && a[i] > a[j]
func kInversePairs(n int, k int) int {
	var mod int64 = 1000000007
	dp := make([][]int64, k+1)
	dp[0] = make([]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < len(dp); i++ {
		dp[i] = make([]int64, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i][j-1])%mod
			if i >= j {
				fmt.Println(i, j)
				dp[i][j] -= dp[i-j][j-1]
			}
			//if j > i {
			//	dp[i][j] = dp[i-1][j] + dp[i][j-1]
			//} else {
			//	for x, y := i, j; y > 0 && x >= 0; x, y = x - 1, y - 1 {
			//		dp[i][j] = (dp[i][j] + dp[x][j-1])%mod
			//	}
			//}
		}
	}
	fmt.Println(dp)
	return int(dp[k][n])
}
