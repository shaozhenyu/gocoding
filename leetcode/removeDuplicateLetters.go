package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(removeDuplicateLetters("bcabc"), "abc")
	fmt.Println(removeDuplicateLetters("cbacdcbc"), "acdb")
	fmt.Println(removeDuplicateLetters("thesqtitxyetpxloeevdeqifkz"), "acdb")
}

func removeDuplicateLetters(s string) string {
	fmt.Println(s)
	if len(s) <= 1 {
		return s
	}
	smallestIdx, i := 0, 0
	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'a')]++
	}
	for ; i < len(s); i++ {
		if s[i] < s[smallestIdx] {
			smallestIdx = i
		}
		cnt[int(s[i]-'a')]--
		if cnt[int(s[i]-'a')] == 0 {
			break
		}
	}
	return string(s[smallestIdx]) + removeDuplicateLetters(strings.Replace(s[smallestIdx+1:], string(s[smallestIdx]), "", -1))
}

func removeDuplicateLetters1(s string) string {
	return rdl(s, "")
}

func rdl(s string, ret string) string {
	exists := make(map[byte]bool)
	for i := 0; i < len(ret); i++ {
		exists[ret[i]] = true
	}
	for i := 0; i < len(s); i++ {
		if !exists[s[i]] {
			ret += string(s[i])
			exists[s[i]] = true
		} else {
			for j := 0; j < len(ret); j++ {
				if ret[j] == s[i] {
					return min(rdl(s[i+1:], ret), rdl(s[i+1:], ret[:j]+ret[j+1:]+string(s[i])))
				}
			}
		}
	}
	return ret
}

func min(x, y string) string {
	if x < y {
		return x
	}
	return y
}
