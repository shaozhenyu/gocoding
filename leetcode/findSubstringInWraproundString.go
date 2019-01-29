package main

import "fmt"

func main() {
	p := "abc"
	p = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	fmt.Println(findSubstringInWraproundString(p))
}

func findSubstringInWraproundString(p string) int {
	count := 0
	m := make(map[byte]int)
	pre := -1
	preCount := 0
	for i := 0; i < len(p); i++ {
		v := int(p[i] - 'a')
		if v == (pre + 1)%26 {
			preCount++
			pre++
		} else {
			pre = v
			preCount = 1
		}
		if m[p[i]] > preCount {
			continue
		}
		count += preCount - m[p[i]]
		m[p[i]] = preCount
	}
	return count
}
