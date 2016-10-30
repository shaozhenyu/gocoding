package main

import "fmt"

func main() {
	s := longestPalindrome("abaabbaa")
	fmt.Println("return : ", s)
}

func longestPalindrome(s string) string {
	long := len(s)

	if long == 0 {
		return ""
	}

	for long > 1 {
		for i := 0; i < len(s) - long + 1; i++ {
			fmt.Println(s[i:i+long])
			if isPalindrome(s[i:i+long]) {
				return s[i:i+long]
			}
		}
		long--
	}

	return string(s[0])
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
