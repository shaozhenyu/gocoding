package main

import (
	"fmt"
	"sort"
)

func main() {
	deck := []int{2,3,5,7,11,13,17}
	fmt.Println(deckRevealedIncreasing(deck))
}

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	ret := make([]int, len(deck))
	all := make([]int, len(deck))
	for i := 0; i < len(all); i++ {
		all[i] = i
	}
	i := 0
	for len(all) > 0 {
		ret[all[0]] = deck[i]
		i++
		if len(all) == 1 {
			break
		}
		all = append(all, all[1])
		all = all[2:]
	}
	return ret
}
