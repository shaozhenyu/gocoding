package main

import "fmt"

func main() {
	fmt.Println(maxPorfit(2, []int{1, 2,3,3,2,9}))
}

func maxPorfit(k int, prices []int) int {
	if k >= len(prices)/2 {
		ret := 0
		for i := 1; i < len(prices); i++ {
			if v := prices[i] - prices[i-1]; v > 0 {
				ret += v
			}
		}
		return ret
	}
	buy := make([]int, k + 1)
	sell := make([]int, k + 1)
	for i := 0; i <= k; i++ {
		buy[i] = -2 << 31
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1] - prices[i])
			sell[j] = max(sell[j], buy[j] + prices[i])
		}
	}
	return sell[k]
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

func maxProfit(k int, prices []int) int {
	m := make(map[[2]int]int)
	return mp(k, prices, m)
}

func mp(k int, prices []int, m map[[2]int]int) int {
	if k == 0 {
		return 0
	}
	if len(prices) == 0 {
		return 0
	}
	if v, ok := m[[2]int{k, len(prices)}]; ok {
		return v
	}
	buy := -1
	max := 0
	for i := 0; i < len(prices); i++ {
		if buy == -1 {
			buy = prices[i]
			continue
		}
		if prices[i] < buy {
			buy = prices[i]
		} else if prices[i] > buy {
			for i + 1 < len(prices) && prices[i+1] > prices[i] {
				i++
			}
			t := prices[i] - buy + maxProfit(k-1, prices[i+1:])
			if t > max {
				max = t
			}
		}
	}
	m[[2]int{k, len(prices)}] = max
	return max
}

