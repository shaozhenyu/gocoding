package main

import "fmt"

func main() {
	strs := []string{"abc", "abcc", "ab"}

	str := longestCommonPrefix(strs)

	fmt.Println(str)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	str := strs[0]

	var j int

	for i := 1; i < len(strs); i++ {
		for j = 0; j < len(strs[i]); j++ {
			if j >= len(str) {
				break
			}
			if str[j] != strs[i][j] {
				str = str[:j]
				break
			}
		}

		if j == len(strs[i]) {
			str = strs[i]
		}
	}

	return str
}
