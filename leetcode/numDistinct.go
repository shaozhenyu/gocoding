package main

import "fmt"

func main() {
	// s := "babgbag"
	// t := "bag"
	s := "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
	t := "bcddceeeebecbc"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	dp := make([][]int, len(t))
	for i := 0; i < len(t); i++ {
		dp[i] = make([]int, len(s))
		for j := i; j < len(s); j++ {
			add := 0
			if t[i] == s[j] {
				add = 1
			}
			if i == 0 {
				if j == 0 {
					dp[i][j] = add
				} else {
					dp[i][j] = add + dp[i][j-1]
				}
			} else {
				if j == 0 {
					dp[i][j] = 0
				} else {
					if add == 1 {
						dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
					} else {
						dp[i][j] = dp[i][j-1]
					}
				}
			}
		}
	}
	return dp[len(t)-1][len(s)-1]
}

