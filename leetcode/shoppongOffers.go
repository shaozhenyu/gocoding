package main

import "fmt"

func main() {
	price := []int{2, 5}
	special := [][]int{[]int{3, 0, 5}, []int{1, 2, 10}}
	needs := []int{3, 2}
	fmt.Println(shoppingOffers(price, special, needs))
}
func shoppingOffers(price []int, special [][]int, needs []int) int {
	least := 0
	for i := 0; i < len(needs); i++ {
		least += price[i] * needs[i]
	}
	shopping(price, special, needs, 0, 0, &least)
	return least
}

func shopping(price []int, special [][]int, needs []int, start int, money int, least *int) {
	fmt.Println(start, needs, *least)
	Loop:
	for i := start; i < len(special); i++ {
		for n := 0; n < len(needs); n++ {
			if special[i][n] > needs[n] {
				continue Loop
			}
		}
		left := 0
		for n := 0; n < len(needs); n++ {
			needs[n] -= special[i][n]
			left += needs[n] * price[n]
		}
		money += special[i][len(needs)]
		if money >= *least {
			for n := 0; n < len(needs); n++ {
				needs[n] += special[i][n]
			}
			money -= special[i][len(needs)]
			continue
		}
		if money+left < *least {
			*least = money + left
		}
		for j := i; j < len(special); j++ {
			shopping(price, special, needs, j, money, least)
		}
		for n := 0; n < len(needs); n++ {
			needs[n] += special[i][n]
		}
		money -= special[i][len(needs)]
	}
}
