package main

import "fmt"

func main() {
	s := "dsahjpjauf"
	words := []string{"ahjpjau","ja","ahbwzgqnuk","tnmlanowax"}
	fmt.Println(numMatchingSubseq(s, words))
}

func numMatchingSubseq(S string, words []string) int {
	m := make(map[byte][]int)
	for i := 0; i < len(S); i++ {
		b := S[i]
		if _, ok := m[b]; !ok {
			m[b] = make([]int, 0)
		}
		m[b] = append(m[b], i)
	}

	count := 0
	for i := 0; i < len(words); i++ {
		n := make(map[byte]int)
		start := -1
		j := 0
		for ; j < len(words[i]); j++ {
			mIdx, ok := m[words[i][j]]
			if !ok {
				break
			}
			t, ok := n[words[i][j]]
			if !ok {
				t = 0
			}
			for {
				if t >= len(mIdx) {
					t = -1
					break
				}
				idx := mIdx[t]
				if idx > start {
					n[words[i][j]] = t
					t = idx
					break
				}
				t++
			}
			if t < start {
				break
			}
			start = t
		}
		if j == len(words[i]) {
			fmt.Println(words[i])
			count++
		}
	}
	return count
}
