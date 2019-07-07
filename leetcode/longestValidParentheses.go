package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("((())))((()))()"))
	fmt.Println(longestValidParentheses("()(()((())"))
	fmt.Println(longestValidParentheses("("))
}

// ((())))((()))()
func longestValidParentheses(s string) int {
	max := 0
	stack := make([]int, len(s) + 1)
	idx := 0
	preCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			idx++
			stack[idx] = preCount
			preCount = 0
		} else {
			if idx <= 0 {
				preCount = 0
			} else {
				preCount += 2 + stack[idx]
				idx--
			}
		}
		if preCount > max {
			max = preCount
		}
	}
	return max
}
