package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "3+2*2/2 + 7-12*1"
	fmt.Println(calculate(s))
}

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return 0
	}
	start := 0
	ret := 0
	vals := make([]int, 0)
	exps := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			i++
		}
		v, _ := strconv.Atoi(s[start:i])
		for i < len(s) && (s[i] == '*' || s[i] == '/') {
			exp := s[i]
			i++
			start = i
			for i < len(s) && (s[i] >= '0' && s[i] <= '9') {
				i++
			}
			v1, _ := strconv.Atoi(s[start:i])
			if exp == '*' {
				v *= v1
			} else {
				v /= v1
			}
		}
		vals = append(vals, v)
		if i < len(s) {
			exps = append(exps, s[i])
		}
		start = i+1
	}
	if len(vals) == 1 {
		return vals[0]
	}
	ret = vals[0]
	for i := 1; i < len(vals); i++ {
		if exps[i-1] == '+' {
			ret += vals[i]
		} else {
			ret -= vals[i]
		}
	}
	return ret
}
