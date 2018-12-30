package main

import "fmt"

func main() {
	s := "()())()"
	fmt.Println(removeInvalidParentheses(s))
}

// "()())()"
// ["()()()", "(())()"]
func removeInvalidParentheses(s string) []string {
	visited := make(map[string]bool)
	ret := make(map[string]bool)
	removeRight(s, visited, ret)
	res := make([]string, 0, len(ret))
	for key := range ret {
		res = append(res, key)
	}
	return res
}

func removeRight(s string, visited, ret map[string]bool) {
	match := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			match++
		} else if s[i] == ')' {
			match--
		}
		if match < 0 {
			for j := 0; j <= i; j++ {
				if s[j] == ')' {
					s1 := s[:j] + s[j+1:]
					if visited[s1] {
						continue
					}
					visited[s1] = true
					removeRight(s1, visited, ret)
				}
			}
			return
		}
	}
	if match == 0 {
		ret[s] = true
		return
	}
	removeLeft(s, visited, ret)
}

func removeLeft(s string, visited, ret map[string]bool) {
	match := 0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '(' {
			match--
		} else if s[i] == ')' {
			match++
		}
		if match < 0 {
			for j := len(s)-1; j >= i; j-- {
				if s[j] == '(' {
					s1 := s[:j] + s[j+1:]
					if visited[s1] {
						continue
					}
					visited[s1] = true
					removeLeft(s1, visited, ret)
				}
			}
			return
		}
	}
	// if match == 0 {
		ret[s] = true
	// }
}
