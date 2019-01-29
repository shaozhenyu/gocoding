package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strWithout3a3b(33, 50))
}

func strWithout3a3b(A int, B int) string {
	if A == B {
		return strings.Repeat("ab", A)
	}
	a := byte('a')
	b := byte('b')
	if A < B {
		a, b = b, a
		A, B = B, A
	}
	twoAB := A - B
	ret := ""
	if B - twoAB < 0 {
		twoAB = B
		ret = strings.Repeat(string([]byte{a, a, b}), twoAB) + strings.Repeat(string([]byte{a}), A - 2 * B)
	} else {
		ret = strings.Repeat(string([]byte{a, a, b}), twoAB) + strings.Repeat(string([]byte{a, b}), B - twoAB)
	}
	fmt.Println(len(ret), A + B)
	return ret
}
