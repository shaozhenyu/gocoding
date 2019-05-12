package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("ececabbacec"))
}

func validPalindrome(s string) bool {
	vp := func(s string) bool {
		fmt.Println(s)
		for i := 0; i < len(s); i++ {
			if s[i] != s[len(s) - i - 1] {
				return false
			}
		}
		return true
	}
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s) - i - 1] {
			if len(s) - i - 2 >= i && s[i] == s[len(s) - i - 2] {
				if vp(s[i:len(s) - i - 1]) {
					return true
				}
			}
			if i + 1 <= len(s) - i - 1 && s[i+1] == s[len(s) - i - 1] {
				return vp(s[i+1:len(s) - i])
			}
			return false
		}
	}
	return true
}
