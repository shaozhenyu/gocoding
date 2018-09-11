package main

import "fmt"

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	s := make([][]int, len(s1) + 1)
	s[0] = make([]int, len(s2) + 1)
	for i := 1; i <= len(s1); i++ {
		s[i] = make([]int, len(s2) + 1)
		s[i][0] = s[i-1][0] + int(s1[i-1])
	}
	for i := 1; i <= len(s2); i++ {
		s[0][i] = s[0][i-1] + int(s2[i-1])
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			add := int(s1[i-1]) + int(s2[j-1])
			if s1[i-1] == s2[j-1] {
				add = 0
			}
			s[i][j] = min(s[i-1][j] + int(s1[i-1]), s[i][j-1] + int(s2[j-1]))
			s[i][j] = min(s[i][j], s[i-1][j-1] + add)
		}
	}
	fmt.Println(s)
	return s[len(s1)][len(s2)]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

