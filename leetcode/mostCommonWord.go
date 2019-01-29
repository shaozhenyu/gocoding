package main

import (
	"fmt"
	"strings"
)

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	fmt.Println(mostCommonWord(paragraph, banned))
}

func mostCommonWord(paragraph string, banned []string) string {
	mb := make(map[string]bool)
	for i := 0; i < len(banned); i++ {
		mb[banned[i]] = true
	}
	mc := make(map[string]int)
	max := 0
	ret := ""
	for i := 0; i < len(paragraph); i++ {
		for i < len(paragraph) && !isAlphabet(paragraph[i]) {
			i++
		}
		start := i
		for i < len(paragraph) && isAlphabet(paragraph[i]) {
			i++
		}
		v := paragraph[start:i]
		v = strings.ToLower(v)
		if len(v) == 0 || mb[v] {
			continue
		}
		mc[v]++
		if max < mc[v] {
			max = mc[v]
			ret = v
		}
	}
	return ret
}

func isAlphabet(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

