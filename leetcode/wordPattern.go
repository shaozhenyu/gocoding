package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abba"
	str := "aa bb bb cc"
	fmt.Println(wordPattern(pattern, str))
}

func wordPattern(pattern string, str string) bool {
	t := strings.Split(str, " ")
	if len(t) != len(pattern) {
		return false
	}
	exist := make(map[string]byte)
	m := make(map[byte]string)
	for i := 0; i < len(pattern); i++ {
		if _, ok := exist[t[i]]; !ok {
			exist[t[i]] = pattern[i]
		}
		if exist[t[i]] != pattern[i] {
			return false
		}
		if _, ok := m[pattern[i]]; !ok {
			m[pattern[i]] = t[i]
		}
		if m[pattern[i]] != t[i] {
			return false
		}
	}
	return true
}
