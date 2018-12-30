package main

import (
	"fmt"
	"strings"
)

func main() {
	wordlist := []string{"KiTe", "kite", "hare", "Hare"}
	queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
	fmt.Println(spellchecker(wordlist, queries))
}

func spellchecker(wordlist []string, queries []string) []string {
	list1 := make(map[string]bool)
	list2 := make(map[string]string)
	list3 := make(map[string]string)
	aeiou := map[byte]bool{'a': true, 'A': true, 'e': true, 'E': true, 'i': true, 'I': true, 'o': true, 'O': true, 'u': true, 'U': true}
	for i := 0; i < len(wordlist); i++ {
		list1[wordlist[i]] = true
		toLower := strings.ToLower(wordlist[i])
		if _, ok := list2[toLower]; !ok {
			list2[toLower] = wordlist[i]
		}
		b := []byte(wordlist[i])
		for j := 0; j < len(b); j++ {
			if aeiou[b[j]] {
				b[j] = '*'
			}
		}
		toLower = strings.ToLower(string(b))
		if _, ok := list3[toLower]; !ok {
			list3[toLower] = wordlist[i]
		}
	}
	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)

	for i := 0; i < len(queries); i++ {
		if list1[queries[i]] {
			continue
		}
		if v, ok := list2[strings.ToLower(queries[i])]; ok {
			queries[i] = v
			continue
		}
		b := []byte(queries[i])
		for j := 0; j < len(b); j++ {
			if aeiou[b[j]] {
				b[j] = '*'
			}
		}
		if v, ok := list3[strings.ToLower(string(b))]; ok {
			queries[i] = v
			continue
		}
		fmt.Println(queries[i])
		queries[i] = ""
	}
	return queries
}
