package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "LeetCode@LeetCode.com"
	fmt.Println(maskPII(s))
	fmt.Println(maskPII("1(234)567-890"))
}

func maskPII(S string) string {
    b := []byte(S)
    if bytes.ContainsRune(b, '@') {
		return maskEmail(b)
	}
	return maskPhone(b)
}

func maskEmail(b []byte) string {
	idx := bytes.LastIndexByte(b, byte('@'))
	pre := []byte("*******")
	pre[0] = b[0]
	pre[6] = b[len(b)-1]
	return strings.ToLower(string(append(pre, b[idx:]...)))
}

func maskPhone(b []byte) string {
	suf := []byte("***-***-")
	nb := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] >= '0' && b[i] <= '9' {
			nb = append(nb, b[i])
		}
	}
	pre := ""
	if len(nb) > 10 {
		p := len(nb) - 10
		pre = "+"
		for p > 0 {
			pre += "*"
		}
		pre += "-"
	}
	suf = append(suf, nb[len(nb)-4:]...)
	return pre + string(suf)
}
