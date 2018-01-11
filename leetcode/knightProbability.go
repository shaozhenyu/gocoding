package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(knightProbability(3, 3, 0, 0))
	fmt.Println(knightProbability(8, 30, 6, 4))
}

func knightProbability(N int, K int, r int, c int) float64 {
	move := [][]int{[]int{-1, -2}, []int{-2, -1}, []int{-1, 2}, []int{-2, 1}, []int{1, -2}, []int{2, -1}, []int{1, 2}, []int{2, 1}}
	dp := make([][]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]float64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = 1
		}
	}

	var x, y int
	for k := 0; k < K; k++ {
		newDp := make([][]float64, N)
		for i := 0; i < N; i++ {
			newDp[i] = make([]float64, N)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for m := 0; m < 8; m++ {
					x = i + move[m][0]
					y = j + move[m][1]
					if isOk(N, x, y) {
						newDp[i][j] += dp[x][y]
					}
				}
			}
		}
		dp = newDp
	}
	return float64(dp[r][c]) / math.Pow(8, float64(K))
}

func isOk(N, r, c int) bool {
	if r < 0 || r >= N || c < 0 || c >= N {
		return false
	}
	return true
}

func knightProbability1(N int, K int, r int, c int) float64 {
	if r < 0 || r >= N || c < 0 || c >= N {
		return 0
	}

	if r+2 < N && r-2 >= 0 && c+2 < N && c-2 >= 0 && K == 1 {
		return 1
	}

	if K == 0 {
		return 1
	}

	return knightProbability(N, K-1, r-2, c-1)/8 + knightProbability(N, K-1, r-1, c-2)/8 + knightProbability(N, K-1, r+1, c-2)/8 + knightProbability(N, K-1, r+2, c-1)/8 + knightProbability(N, K-1, r-1, c+2)/8 + knightProbability(N, K-1, r-2, c+1)/8 + knightProbability(N, K-1, r+2, c+1)/8 + knightProbability(N, K-1, r+1, c+2)/8
}
