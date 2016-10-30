package main

import "fmt"

func main() {
	needle := "bcabc"
	haystack := "abcabc"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {

	len1 := len(haystack)
	len2 := len(needle)

	if (len1 == 0 || len2 == 0) {
		return -1
	}

	if (len1 < len2) {
		return -1
	}

	var j int
	for i := 0; i < len1; i++ {
		if (i + len2) > len1 {
			return -1
		}
		m := i
		for j = 0; j < len2; j++ {
			//fmt.Println(string(haystack[m]), string(needle[j]))
			if haystack[m] != needle[j] {
				break
			}
			m++
		}
		if j == len2 {
			return i+1
		}
	}

	return -1
}
