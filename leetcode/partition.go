package main

import "fmt"

func main() {
	fmt.Println(partition("aaba"))
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	vals := [][]string{}
	for i := 0; i < len(s); i++ {
		if !check(s[:i+1]) {
			continue
		}
		if i == len(s) - 1 {
			vals = append(vals, []string{s[:i+1]})
			return vals
		}
		val := partition(s[i+1:])
		for _, v := range val {
			s := append([]string{s[:i+1]}, v...)
			vals = append(vals, s)
		}
	}
	return vals
}

func check(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
