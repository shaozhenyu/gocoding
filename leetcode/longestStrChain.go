package main

import "fmt"

func main() {
	fmt.Println(longestStrChain([]string{"a","b","ba","bca","bda","bdca"}))
}

func longestStrChain(words []string) int {
	if len(words) == 0 {
		return 0
	}
	start := 2 << 31
	size := make(map[int]map[string]int)
	for i := 0; i < len(words); i++ {
		l := len(words[i])
		if l < start {
			start = l
		}
		if _, ok := size[l]; !ok {
			size[l] = make(map[string]int)
		}
		size[l][words[i]] = 1
	}
	start++
	max := 1
	for {
		if _, ok := size[start]; !ok {
			break
		}
		for k := range size[start] {
			t := size[start][k]
			for i := 0; i < len(k); i++ {
				if num, ok := size[start-1][k[:i] + k[i+1:]]; ok {
					tt := t + num
					fmt.Println(t, num, k, k[:i] + k[i+1:])
					if tt > size[start][k] {
						size[start][k] = tt
					}
				}
			}
			if size[start][k] > max {
				max = size[start][k]
			}
		}
		start++
	}
	fmt.Println(len(words), words[0], words[len(words)-1])
	return max
}
