package main

import "fmt"

func main() {
	s := "DIDDI"
	fmt.Println(numPermsDISequence(s))
}

// "D" 减少 "I" 增加
func numPermsDISequence(S string) int {
	n := len(S)
	m := make(map[int]bool)
	for i := 0; i <= n; i++ {
		m[i] = true
	}
	count := 0
	for i := 0; i <= n; i++ {
		delete(m, i)
		count += numPerms(S, m, i, n)
		m[i] = true
	}
	return count
}

func numPerms(S string, m map[int]bool, now, n int) int {
	if len(S) == 0 {
		if len(m) == 0 {
			return 1
		}
		return 0
	}
	count := 0
	if S[0] == 'D' {
		for i := now - 1; i >= 0; i-- {
			if m[i] {
				delete(m, i)
				count += numPerms(S[1:], m, i, n)
				m[i] = true
			}
		}
	} else {
		for i := now + 1; i <= n; i++ {
			if m[i] {
				delete(m, i)
				count += numPerms(S[1:], m, i, n)
				m[i] = true
			}
		}
	}
	return count
}
