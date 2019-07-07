package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("aa bb   cc"))
}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	ret := ""
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] != "" {
			ret += ss[i] + " "
		}
	}
	if len(ret) == 0 {
		return ret
	}
	return ret[:len(ret)-1]
}
