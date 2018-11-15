package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcabcabc"
	fmt.Println(repeatedSubstringPattern(s))
}

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		if strings.Replace(s, s[:i], "", -1) == "" {
			return true
		}
	}
	return false
}
