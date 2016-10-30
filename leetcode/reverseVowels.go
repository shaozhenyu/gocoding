package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	vowels := map[string]int{
		"a": 1,
		"e": 1,
		"i": 1,
		"o": 1,
		"u": 1,
	}

	i := 0
	j := len(s) - 1
	str := make([]string, len(s))

	for i <= j {

		_, ok1 := vowels[string(s[i])]
		_, ok2 := vowels[string(s[j])]

		if ok1 && ok2 {
			str[i] = string(s[j])
			str[j] = string(s[i])
			i++
			j--
		} else if ok1 && !ok2 {
			str[j] = string(s[j])
			j--
		} else if !ok1 && ok2 {
			str[i] = string(s[i])
			i++
		} else {
			str[i] = string(s[i])
			str[j] = string(s[j])
			i++
			j--
		}
	}
	
	ret := ""
	for _, v := range str {
		ret += v
	}
	
	return ret
}
