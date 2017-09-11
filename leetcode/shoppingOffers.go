package main

import (
	"fmt"
)

func main() {

	price := []int{9, 6, 1, 5, 3, 4}
	special := [][]int{{1, 2, 2, 1, 0, 4, 14}, {6, 3, 4, 0, 0, 1, 16}}
	needs := []int{6, 6, 6, 1, 6, 6}
	fmt.Println(shoppingOffers(price, special, needs))
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	ret := 99999
	lenS := len(special)

	for i := 0; i < lenS; i++ {
		ok := true
		for j := 0; j < len(needs); j++ {
			needs[j] -= special[i][j]
			if needs[j] < 0 {
				ok = false
			}
		}

		if ok {
			ret = min(ret, shoppingOffers(price, special, needs)+special[i][len(needs)])
		}

		for j := 0; j < len(needs); j++ {
			needs[j] += special[i][j]
		}
	}

	noSpecial := 0
	for i := 0; i < len(needs); i++ {
		noSpecial += price[i] * needs[i]
	}

	return min(ret, noSpecial)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
