package main

import (
	"fmt"
	"math"
)

func main() {
	m, n, N, i, j := 1, 3, 3, 0, 1
	fmt.Println(findPaths(m, n, N, i, j))
}

var step [][]int = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}

func findPaths(m int, n int, N int, i int, j int) int {
	mod := int(math.Pow(10, 9)) + 7
	dp := make([][]int, m)
	for ii := 0; ii < m; ii++ {
		dp[ii] = make([]int, n)
		for jj := 0; jj < n; jj++ {
			if ii == 0 {
				dp[ii][jj]++
			}
			if ii == m - 1 {
				dp[ii][jj]++
			}
			if jj == 0 {
				dp[ii][jj]++
			}
			if jj == n - 1 {
				dp[ii][jj]++
			}
		}
	}
	ret := 0
	for N > 0 {
		N--
		ret = (ret + dp[i][j])%mod
		tmp := make([][]int, m)
		for ii := 0; ii < m; ii++ {
			tmp[ii] = make([]int, n)
			for jj := 0; jj < n; jj++ {
				for _, s := range step {
					x := ii + s[0]
					y := jj + s[1]
					if x >= 0 && x < m && y >= 0 && y < n {
						tmp[ii][jj] = (tmp[ii][jj] + dp[x][y])%mod
					}
				}
			}
		}
		dp = tmp
	}
	return ret
}

