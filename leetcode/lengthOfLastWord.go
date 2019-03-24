package main

import (
	"fmt"
)

func main() {
	s := "hello world  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	i := len(s) - 1
	for i >= 0 {
		if s[i] == byte(' ') {
			i--
		} else {
			break
		}
	}
	for ; i >= 0; i-- {
		if s[i] != byte(' ') {
			count++
		} else {
			break
		}
	}
	return count
}
