package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("ababccabc"))
}

func isValid(S string) bool {
	for len(S) > 0 {
		old := S
		S = strings.Replace(old, "abc", "", -1)
		if old == S {
			return false
		}
	}
	return true
}
