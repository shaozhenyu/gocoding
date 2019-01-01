package main

import "fmt"

func main() {
	s := "cdaabbac"
	fmt.Println(longestSubstring(s, 2))
}

// cdaabbacd
func longestSubstring(s string, k int) int {
	if len(s) == 0 || k == 1 {
		return len(s)
	}
	bk := make(map[byte]int)
	lk := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if bk[s[i]] >= k {
			bk[s[i]]++
		} else {
			lk[s[i]]++
			if lk[s[i]] >= k {
				bk[s[i]] = lk[s[i]]
				delete(lk, s[i])
			}
		}
	}
	if len(lk) == 0 {
		return len(s)
	}
	start := 0
	ret := 0
	for i := 0; i < len(s); i++ {
		if _, ok := lk[s[i]]; ok {
			ret = max(ret, longestSubstring(s[start:i], k))
			start = i + 1
		}
	}
	return max(ret, longestSubstring(s[start:], k))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
