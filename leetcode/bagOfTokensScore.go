package main

import (
	"fmt"
	"sort"
)

func main() {
	tokens := []int{100,200,300,400}
	P := 200
	fmt.Println(bagOfTokensScore(tokens, P))
}

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	score, maxScore := 0, 0
	for i, j := 0, len(tokens) - 1; i <= j; {
		if P >= tokens[i] {
			score++
			P -= tokens[i]
			i++
			if score > maxScore {
				maxScore = score
			}
		} else {
			if score == 0 {
				break
			}
			score--
			P += tokens[j]
			j--
		}
	}
	return maxScore
}
