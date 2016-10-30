package main

import "fmt"

func main() {
	if canConstruct("aa", "aba") {
		fmt.Println("11111")
	}
}

func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	r := make([]int, 26)
	m := make([]int, 26)

	var i int

	for i = 0; i < len(ransomNote); i++ {
		r[ransomNote[i]-97]++
		m[magazine[i]-97]++
	}

	for i < len(magazine) {
		m[magazine[i]-97]++
		i++
	}

	for i = 0; i < 26; i++ {
		if r[i] > m[i] {
			return false
		}
	}

	return true
}
