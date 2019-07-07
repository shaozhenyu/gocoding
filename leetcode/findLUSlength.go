package main

import "fmt"

func main() {
	fmt.Println(findLUSlength([]string{"aa", "bb", "aa"}))
}

func findLUSlength(strs []string) int {
	mStrs := [11]map[string]int{}
	for i := 0; i <= 10; i++ {
		mStrs[i] = make(map[string]int, 0)
	}
	for _, s := range strs {
		mStrs[len(s)][s]++
	}
	for i := 10; i > 0; i-- {
		for str, count := range mStrs[i] {
			if count == 1 {
				return len(str)
			}
			for j := 0; j < len(str); j++ {
				mStrs[i-1][str[:j] + str[j+1:]] += 2
			}
		}
	}
	return -1
}
