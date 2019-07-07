package main

import (
	"fmt"
)

func main() {
	for i := 3; i <= 3; i++ {
		fmt.Println(i, baseNeg2(i))
	}
}

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}
	ret := ""
	for N != 0 {
		mod := 0
		if N < 0 && N%(-2) != 0 {
			mod = 1
			N = N/(-2) + 1
		} else {
			mod = N%(-2)
			N = N/(-2)
		}
		ret = fmt.Sprintf("%d", mod) + ret
	}
	return ret
}
