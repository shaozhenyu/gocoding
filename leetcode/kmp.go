package main

import "fmt"

func main() {
	fmt.Println(kmpSearch("aaaabddcbcd", "bc"))
}

func kmpSearch(str, pat string) int {
	dfa := DFA(pat)
	i, j := 0, 0
	for i = 0; i < len(str) && j < len(pat); i++ {
		j = dfa[charAt(str[i])][j]
	}
	if j == len(pat) {
		return i - j
	}
	return -1
}

func DFA(pat string) [][]int {
	n := len(pat)
	R := 256
	dfa := make([][]int, R)
	for i := 0; i < R; i++ {
		dfa[i] = make([]int, n)
	}
	dfa[charAt(pat[0])][0] = 1
	for x, j := 0, 1; j < n; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x]
		}
		dfa[charAt(pat[j])][j] = j + 1
		x = dfa[charAt(pat[j])][x]
	}
	return dfa
}

func charAt(x byte) int {
	return int(x - 'a')
}
