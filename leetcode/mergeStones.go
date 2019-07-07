package main

import "fmt"

func main() {
	stones := []int{3,2,4,1}
	K := 2
	fmt.Println(mergeStones(stones, K), 20)
}

func mergeStones(stones []int, K int) int {
	dp := make([][][]int, len(stones))
	sum := make([]int, len(stones))
	tmp := 0
	if (len(stones) - 1)%(K-1) != 0 {
		return -1
	}
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, len(stones))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, K+1)
			for k := 0; k <= K; k++ {
				dp[i][j][k] = 2 << 31
			}
		}
		dp[i][i][1] = 0
		sum[i] = tmp + stones[i]
		tmp += stones[i]
	}
	for length := 1; length < len(stones); length++ {
		for left := 0; left + length < len(stones); left++ {
			right := left + length
			for mid := left; mid < right; mid++ {
				for k := 2; k <= K; k++ {
					dp[left][right][k] = min(dp[left][right][k], dp[left][mid][k-1] + dp[mid+1][right][1])
				}
			}
			dp[left][right][1] = dp[left][right][K] + (sum[right] - sum[left] + stones[left])
		}
	}
	fmt.Println(dp)
	return dp[0][len(stones)-1][1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
