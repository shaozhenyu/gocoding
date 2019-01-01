package main

import (
	"fmt"
	"strconv"
)

func main() {
	chars := []byte("aabbcca")
	fmt.Println(compress(chars))
}

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	count := 0
	same := 1
	idx := 0
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			same++
		}
		if chars[i] != chars[i-1] || i == len(chars) - 1 {
			chars[idx] = chars[i-1]
			idx++
			count++
			s := []int{}
			for same > 1 {
				s = append(s, same%10)
				same /= 10
				count++
			}
			for j := len(s) - 1; j >= 0; j-- {
				chars[idx] = []byte(strconv.Itoa(s[j]))[0]
			}
		}
	}
	return count
}
