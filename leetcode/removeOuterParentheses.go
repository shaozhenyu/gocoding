package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("()(())"))
}

func removeOuterParentheses(S string) string {
	start, count := 0, 0
	ret := ""
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			count++
		} else {
			count--
			if count == 0 {
				ret += S[start+1:i]
				start = i + 1
			}
		}
	}
	return ret
}
