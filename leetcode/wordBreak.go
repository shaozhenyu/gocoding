package main

import (
	"fmt"
)

func main() {
	s := "abcbcaabcd"
	wordDict := []string{"abc", "cbc", "ab", "a", "cd"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	m := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}
	return checkPre(s, m, make(map[string]bool))
}

func checkPre(s string, m map[string]bool, checked map[string]bool) bool {
	if s == "" {
		return true
	}
	if checked[s] {
		return false
	}
	for i := 0; i < len(s); i++ {
		if m[s[:i+1]] && checkPre(s[i+1:], m, checked) {
			return true
		}
	}
	checked[s] = true
	return false
}
