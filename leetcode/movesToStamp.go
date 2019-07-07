package main

import (
	"bytes"
	"fmt"
)

func main() {
	stamp := "abc"
	target := "ababc"
	stamp = "abca"
	target = "aabcaca"
	fmt.Println(movesToStamp(stamp, target))
}

func movesToStamp(stamp string, target string) []int {
	bTarget := make([]byte, len(target))
	bStamp := []byte(stamp)
	for i := 0; i < len(target); i++ {
		bTarget[i] = '?'
	}
	now := []byte(target)
	ret := []int{}
	visited := make(map[int]bool)
	var flag = true
	for flag {
		flag = false
		if bytes.Compare(now, bTarget) == 0 {
			for i := 0; i < len(ret)/2; i++ {
				ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
			}
			return ret
		}
		for i := 0; i <= len(bTarget) - len(stamp); i++ {
			if visited[i] {
				continue
			}
			if compare(bStamp, now[i:i+len(stamp)]) {
				ret = append(ret, i)
				for j := i; j < i + len(stamp); j++ {
					now[j] = '?'
				}
				visited[i] = true
				flag = true
				break
			}
		}
	}
	return []int{}
}

func compare(bs, bt []byte) bool {
	if bytes.Compare(bs, bt) == 0 {
		return true
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] != bt[i] && bt[i] != '?' {
			return false
		}
	}
	return true
}
