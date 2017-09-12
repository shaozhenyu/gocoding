package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubstrings("abccbasddnsakdfbhfqew"))
}

func countSubstrings(s string) int {
	lenS := len(s)
	ret := lenS

	for i := 0; i < lenS - 1; i++ {
		for j := i + 1; j < lenS; j++ {
			if s[i] == s[j] {
				m := i + 1
				n := j - 1
				for m < n {
					if s[m] != s[n] {
						break
					}
					m++
					n--
				}
				if m >= n {
					ret++
				}
				fmt.Println(ret)
			}
		}
	}

	return ret
}
