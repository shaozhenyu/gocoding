package main

import "fmt"

func main() {
	s1 := "abcdeb"
	s2 := "abcdbeadccc"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	sm1 := make(map[byte]int)
	sm2 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		sm1[s1[i]]++
	}
	start := 0
	for i := 0; i < len(s2); i++ {
		if c, ok := sm1[s2[i]]; ok {
			if sm2[s2[i]] == c {
				j := start
				for ; s2[j] != s2[i]; j++ {
					sm2[s2[j]]--
				}
				start = j + 1
			} else {
				sm2[s2[i]]++
				if i - start + 1 == len(s1) {
					return true
				}
			}
		} else {
			sm2 = make(map[byte]int)
			start = i + 1
		}
	}
	return false
}
