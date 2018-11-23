package main

import "fmt"

func main() {
	s := "ababa"
	p := "ab"
	s = "abaacbabc"
	p = "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	ret := []int{}
	match := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		match[p[i]]++
	}

	target := make(map[byte]int)
	start, end := 0, 0
	for start < len(s) && end < len(s) {
		fmt.Println(start, end)
		if _, ok := match[s[end]]; !ok {
			start = end + 1
			for start < len(s) && match[s[start]] == 0 {
				start++
			}
			if start >= len(s) {
				break
			}
			target = make(map[byte]int)
			end = start
			continue
		}
		target[s[end]]++
		if end - start + 1 == len(p) {
			end++
			if check(match, target) {
				ret = append(ret, start)
				for end < len(s) && s[start] == s[end] {
					start++
					end++
					ret = append(ret, start)
				}
			}
			target[s[start]]--
			if start >= len(s) {
				break
			}
			if target[s[start]] == 0 {
				delete(target, s[start])
			}
			start++
		} else {
			end++
		}
	}
	return ret
}

func check(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
