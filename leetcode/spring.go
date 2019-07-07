package main

import "fmt"

func main() {
	fmt.Println(fn("abc", "abcbc"), 2)
	fmt.Println(fn("xyz", "xzyxz"), 3)
	fmt.Println(fn("abc", "acdbc"), -1)
	fmt.Println(fn("abcabc", "abcb"), 1)
}

func fn(source, target string) int {
	ms := make(map[byte][]int)
	for i := 0; i < len(source); i++ {
		if _, ok := ms[source[i]]; !ok {
			ms[source[i]] = make([]int, 0)
		}
		ms[source[i]] = append(ms[source[i]], i)
	}
	for i := 0; i < len(target); i++ {
		if _, ok := ms[target[i]]; !ok {
			return -1
		}
	}
	return f(target, ms)
}

func f(target string, ms map[byte][]int) int {
	if len(target) == 0 {
		return 0
	}
	start := -1
	for i := 0; i < len(target); i++ {
		idxs := ms[target[i]]
		flag := false
		for _, idx := range idxs {
			if idx > start {
				start = idx
				flag = true
				break
			}
		}
		if !flag {
			return 1 + f(target[i:], ms)
		}
	}
	return 1
}
