package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestRepeatingSubstring("aaaaa"))
}

func longestRepeatingSubstring(S string) int {
    maxLen := 0
    for i := 0; i < len(S); i++ {
        j := i + maxLen
        for ; j < len(S); j++ {
			fmt.Println(S[i:j+1], strings.Count(S, S[i:j+1]))
            if strings.Count(S[i+1:], S[i:j+1]) > 1 {
                maxLen = j - i + 1
            }
        }
    }
    return maxLen
}
