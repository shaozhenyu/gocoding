package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(licenseKeyFormatting("2-4A0r7-4k", 3))
}

func licenseKeyFormatting(S string, K int) string {
	s := strings.Replace(S, "-", "", -1)
	s = strings.ToUpper(s)
	ret := make([]byte, len(s)+len(s)/K)
	mod := 0
	j := len(ret) - 1
	for i := len(s) - 1; i >= 0; i-- {
		ret[j] = s[i]
		j--
		mod++
		if mod%K == 0 {
			ret[j] = '-'
			j--
		}
	}
	return string(bytes.TrimLeft(ret, '-'))
}
