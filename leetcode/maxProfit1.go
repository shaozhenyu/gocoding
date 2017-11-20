package main

import (
	"fmt"
)

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
}

func maxProfit(prices []int, fee int) int {
	profit := 0
	pre := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-pre-fee > 0 {
			bigger := prices[i]
			var j int
			for j = i + 1; j < len(prices); j++ {
				if prices[j]+fee <= bigger {
					break
				}
				if prices[j] < bigger {
					continue
				}
				bigger = prices[j]
			}
			profit += bigger - pre - fee
			if j == len(prices) {
				break
			}
			pre = prices[j]
		} else {
			if prices[i] < pre {
				pre = prices[i]
			}
		}
	}
	return profit
}
