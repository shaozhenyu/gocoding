package main

import "fmt"

func main() {
	fmt.Println(getMoneyAmount(10))
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := 0; i <= n + 1; i++ {
		dp[i] = make([]int, n+2)
	}
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 1; j-- {
			dp[j][i] = 1 << 32 - 1
			for k := j; k <= i; k++ {
				dp[j][i] = min(dp[j][i], k + max(dp[j][k-1], dp[k+1][i]))
			}
		}
	}
	fmt.Println(dp)
	return dp[1][n]
}

func gma(start, end int) int {
	fmt.Println(start, end)
	if start >= end {
		return 0
	}
	if end - start == 1 {
		return start
	}
	mid := start + (end - start)/2
	if (end - start)%2 == 1 {
		mid++
	}
	return mid + max(gma(mid+1, end), gma(start, mid-1))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
