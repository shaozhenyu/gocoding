package main

import "fmt"

func main() {
	S := "ADOBECODEBANC"
	T := "ABC"
	S = "abc"
	T = "ac"
	S = "acbbaca"
	T = "aba"
	fmt.Println(minWindow(S, T))
}

func minWindow(s string, t string) string {
	if s == t {
		return s
	}
	m := make(map[byte]int)
	n := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	now, start := 0, 0
	success := len(m)
	min := len(s) + 1
	ret := ""
	checkMin := func(start, end int) {
		size := end + 1 - start
		if size < min {
			min = size
			ret = s[start:end+1]
		}
	}
	for i := start; i < len(s); i++ {
		if m[s[i]] == 0 {
			continue
		}
		n[s[i]]++
		if n[s[i]] == m[s[i]] {
			now++
			if now < success {
				continue
			}
			checkMin(start, i)
			for start < i {
				if m[s[start]] == 0 {
					start++
					checkMin(start, i)
				} else {
					n[s[start]]--
					if n[s[start]] < m[s[start]] {
						now--
						start++
						break
					} else {
						start++
						checkMin(start, i)
					}
				}
			}
		}
	}
	return ret
}
