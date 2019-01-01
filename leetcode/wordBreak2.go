package main

import "fmt"

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	wordDict = []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}
	_, words := wordBreak1(s, m, make(map[string][][]string))
	ret := make([]string, 0)
	for _, word := range words {
		s := ""
		for i := 0; i < len(word); i++ {
			s += word[i] + " "
		}
		ret = append(ret, s[:len(s)-1])
	}
	return ret
}

func wordBreak1(s string, m map[string]bool, exist map[string][][]string) (bool, [][]string) {
	if _, ok := exist[s]; ok {
		fmt.Println("xx:", s, exist[s])
		return true, exist[s]
	}
	ret := make([][]string, 0)
	for i := 0; i < len(s); i++ {
		if m[s[:i+1]] {
			if i + 1 == len(s) {
				ret = append(ret, []string{s[:i+1]})
			} else {
				_, t := wordBreak1(s[i+1:], m, exist)
				// exist[s[i+1:]] = t
				for _, t1 := range t {
					ret = append(ret, append([]string{s[:i+1]}, t1...))
				}
			}
		}
	}
	exist[s] = ret
	return true, ret
}
