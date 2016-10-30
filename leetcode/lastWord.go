package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("a aaa "))
}

func lengthOfLastWord(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	n := 0

	for i := length - 1; i >= 0; i-- {

		if s[i] < 65 || (s[i] > 90 && s[i] < 97) || s[i] > 122 {
			if n == 0 {
				continue
			} else {
				if s[i] != 32 {
					return 0
				}
			}
		}

		if s[i] == 32 {
			return n
		} else {
			n++
		}
	}

	return n
}
