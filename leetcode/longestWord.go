package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"b", "applye", "a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println(longestWord(words))
}

func longestWord(words []string) string {
	if len(words) == 0 {
		return ""
	}
	have := make(map[string]bool)
	sort.Strings(words)
	fmt.Println(words)
	ret := ""
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 1 {
			have[words[i]] = true
			if ret == "" {
				ret = words[i]
			}
		} else {
			if have[words[i][:len(words[i])-1]] {
				have[words[i]] = true
				if len(words[i]) > len(ret) {
					ret = words[i]
				}
			}
		}
	}
	return ret
}
