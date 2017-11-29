package main

import (
	"fmt"
)

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3

	strs = []string{"10", "0", "1"}
	m = 1
	n = 1

	strs = []string{"10", "0001", "111001", "1", "0"}
	m = 3
	n = 4
	fmt.Println(findMaxForm(strs, m, n))
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	var one, zero int
	for i := 0; i < len(strs); i++ {
		one, zero = 0, 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				zero++
			} else {
				one++
			}
		}

		fmt.Println("zero, one:", zero, one)
		for mm := m; mm >= zero; mm-- {
			for nn := n; nn >= one; nn-- {
				dp[mm][nn] = max(dp[mm][nn], dp[mm-zero][nn-one]+1)
			}
		}
		fmt.Println(dp)
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
