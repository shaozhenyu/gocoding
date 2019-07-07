package main

import "fmt"

func main() {
	fmt.Println(bitwiseComplement(0))
}

func bitwiseComplement(N int) int {
	s := fmt.Sprintf("%b", N)
	b := 1
	ret := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			ret += b
		}
		b *= 2
	}
	return ret
}
