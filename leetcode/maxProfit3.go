package main

import (
	"fmt"
)

func main() {
	prices := []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	buy1, buy2 := -1 << 32, -1 << 32
	sell1, sell2 := 0, 0
	for i := 0; i < len(prices); i++ {
		buy1 = max(buy1, -1 * prices[i])
		sell1 = max(sell1, prices[i] + buy1)
		buy2 = max(buy2, sell1 - prices[i])
		sell2 = max(sell2, buy2 + prices[i])
	}
	return sell2
}

func maxProfit1(prices []int) int {
	ret := 0
	if len(prices) <= 1 {
		return ret
	}
	for i := 0; i < len(prices); i++ {
		oneMin := prices[0]
		one := 0
		for j := 1; j < i; j++ {
			one = max(one, prices[j] - oneMin)
			oneMin = min(oneMin, prices[j])
		}
		
		twoMin := prices[i]
		two := 0
		for j := i + 1; j < len(prices); j++ {
			two = max(two, prices[j] - twoMin)
			twoMin = min(twoMin, prices[j])
		}
		this := one + two
		ret = max(ret, this)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
