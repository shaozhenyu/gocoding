package main

import (
	"fmt"
	"sort"
)

func main() {
	coins := []int{186,419,83,408}
	amount := 6249
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for i := 1; i < len(dp); i++ {
		dp[i] = 2 << 31
		for j := 0; j < len(coins); j++ {
			if i - coins[j] < 0 {
				break
			}
			dp[i] = min(dp[i], dp[i - coins[j]] + 1)
		}
	}
	if dp[amount] ==  2 << 31 {
		dp[amount] = -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// slow
func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	m := make(map[int]struct{})
	m[amount] = struct{}{}
	for i := 1; len(m) != 0; i++ {
		m1 := make(map[int]struct{})
		for k := range m {
			for j := 0; j < len(coins); j++ {
				v := k - coins[j]
				if v == 0 {
					return i
				}
				if v > 0 {
					m1[v] = struct{}{}
				}
			}
		}
		m = m1
	}
	return -1
}
