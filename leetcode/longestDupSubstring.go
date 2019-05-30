package main

import (
	"fmt"
	// "hash/crc32"
)

func main() {
	fmt.Println(longestDupSubstring("abanana"))
}

func longestDupSubstring(S string) string {
	fmt.Println("s: ", S)
	for length := len(S) - 1; length > 0; length-- {
		v := RabinKarp(S, length)
		if len(v) == length {
			return v
		}
	}
	return ""
}

func RabinKarp(s string, length int) string {
	mod := 1000000000
	r := 10
	m := make(map[int]bool)
	rm := 1
	// now := int(crc32.ChecksumIEEE([]byte(s[:length]))) % mod
	now := 0
	for i := 0; i < length; i++ {
		now = (now * r + (int(s[i]) - 97 + 1)) % mod
		rm = (rm * r) % mod
	}
	m[now] = true
	fmt.Println("start:", s[:length], now)
	for i := 1; i <= len(s) - length; i++ {
		j := i + length - 1
		now = (now + mod - rm * int(s[i-1]) % mod) % mod
		fmt.Println("11111: ", now)
		now = (now * r + (int(s[j]) - 97 + 1)) % mod
		fmt.Println(s[i:j+1], now)
		if m[now] {
			return s[i:j+1]
		}
		m[now] = true
	}
	fmt.Println(m)
	return ""
}

