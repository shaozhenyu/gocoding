package main

import "fmt"

func main() {
	fmt.Println(canIWin(10, 11), false)
	fmt.Println(canIWin(10, 40), false)
	fmt.Println(canIWin(5, 50), false)
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	sum := 0
	for i := 1; i <= maxChoosableInteger; i++ {
		sum += i
	}
	if sum < desiredTotal {
		return false
	}
	used := make([]bool, maxChoosableInteger + 1)
	dp := make(map[int]bool)
	return canIWin1(used, 0, desiredTotal, dp, 0)
}

func canIWin1(used []bool, sum int, desiredTotal int, dp map[int]bool, dpKey int) bool {
	if _, ok := dp[dpKey]; ok {
		return dp[dpKey]
	}
	for i := len(used) - 1; i >= 1; i-- {
		if !used[i] && (i + sum) >= desiredTotal {
			dp[dpKey] = true
			return true
		}
	}
	for i := len(used) - 1; i >= 1; i-- {
		if !used[i] {
			used[i] = true
			ok := canIWin1(used, sum + i, desiredTotal, dp, (dpKey | (1 << uint(i))))
			if !ok {
				used[i] = false
				dp[dpKey] = true
				return true
			}
			used[i] = false
		}
	}
	dp[dpKey] = false
	return false
}
