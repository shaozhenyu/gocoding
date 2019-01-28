package main

import "fmt"

func main() {
	s := "abcbaa"
	fmt.Println(shortestPalindrome(s))
}

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	i := len(s) - 1
	add := make([]byte, len(s))
	idx := 0
	for ; i > 0; i-- {
		if s[i] == s[0] && isPalindrome(s[1:i]) {
			break
		}
		add[idx] = s[i]
		idx++
	}
	return string(add[:i]) + s
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
