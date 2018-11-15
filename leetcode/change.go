package main

import "fmt"

func main() {
	amount := 10
	coins := []int{1, 2}
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	if amount == 0 || len(coins) == 0 {
		return 0
	}
	dp := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			count := 0
			for k := 0; k * coins[i] <= j; k++ {
				count += dp[i][j - k * coins[i]]
			}
			dp[i+1][j] = count
		}
	}
	fmt.Println(dp)
	return dp[len(coins)][amount]
}

func change_slow(amount int, coins []int) int {
	m := make(map[int]map[int]int)
	for i := 1; i <= amount; i++ {
		m[i] = make(map[int]int)
	}
	return change1(0, amount, coins, m)
}

func change1(idx int, amount int, coins []int, m map[int]map[int]int) int {
	if m[amount][idx] > 0 {
		return m[amount][idx]
	}
	count := 0
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1
	}
	if len(coins) == idx {
		return 0
	}
	now := coins[idx]
	i := 0
	for i * now <= amount {
		count += change1(idx+1, amount - i * now, coins, m)
		i++
	}
	m[amount][idx] = count
	return count
}
