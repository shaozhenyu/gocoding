package main

import "fmt"

func main() {
	s := "abcabcccaa"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	if len(s) <= 1 {
		return len(s)
	}
	maxLen := 1
	m := make(map[byte]int) // ma[value]idx
	start := 0
	m[s[0]] = start
	for i := 0; i < len(s); i++ {
		if m[s[i]] > start {
			start = m[s[i]]
		}
		if i - start + 1 > maxLen {
			maxLen = i - start + 1
		}
		m[s[i]] = i + 1
	}
	return maxLen
}
