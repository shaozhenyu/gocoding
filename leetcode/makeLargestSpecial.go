package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(makeLargestSpecial("1110101011010111000000"))
	fmt.Println("true : 1111110001010010101000")
}

// 11011000
// 10101100
// 10111101000
func makeLargestSpecial(S string) string {
	if len(S) == 0 {
		return ""
	}
	one, zero := 0, 0
	ret := []string{}
	start := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '1' {
			one++
		} else {
			zero++
		}
		if one == zero {
			ret = append(ret, "1" + makeLargestSpecial(S[start+1:i]) + "0")
			start = i+1
			one, zero = 0, 0
		}
	}
	return exchange(ret)
}

func exchange(s []string) string {
	sort.Strings(s)
	ret := ""
	for i := len(s)-1; i >= 0; i-- {
		ret += s[i]
	}
	return ret
}
