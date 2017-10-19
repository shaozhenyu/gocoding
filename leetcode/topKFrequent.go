package main

import (
	"fmt"
	"sort"
	"strings"
)

type Fre struct {
	Word  string
	Times int
}

func main() {
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is", "is", "a", "a", "a", "a"}
	fmt.Println(topKFrequent(words, 4))
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)

	for _, v := range words {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	fre := make([]Fre, len(m))
	var i int = 0
	for k, v := range m {
		fre[i] = Fre{
			Word:  k,
			Times: v,
		}
		i++
	}

	sort.Slice(fre, func(i, j int) bool {
		if fre[i].Times == fre[j].Times {
			com := strings.Compare(fre[i].Word, fre[j].Word)
			if com > 0 {
				return false
			} else {
				return true
			}
		}
		return fre[i].Times > fre[j].Times
	})

	fmt.Println(fre)

	ret := make([]string, k)
	for i = 0; i < k; i++ {
		ret[i] = fre[i].Word
	}

	return ret
}
