package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	S := "0.(52)"
	T := "0.5(25)"
	S = "0.08(9)"
	T = "0.09"
	S = "0.9(9)"
	T = "1."
	S = "0.(0)"
	T = "0"
	fmt.Println(isRationalEqual(S, T))
}

var nine = strings.Repeat("9", 100)

func isRationalEqual(S string, T string) bool {
	s1, s2 := parse(S)
	t1, t2 := parse(T)
	if s2 == strings.Repeat("0", 100) {
		s2 = ""
	}
	if t2 == strings.Repeat("0", 100) {
		t2 = ""
	}
	return (s1 == t1) && (s2 == t2)
}

func parse(str string) (x int, y string) {
	s := strings.Split(str, ".")
	x, _ = strconv.Atoi(s[0])
	if len(s) == 1 || s[1] == "" {
		y = strings.Repeat("0", 100)
		return
	}
	y = s[1]
	if strings.ContainsRune(s[1], '(') {
		start := strings.IndexRune(s[1], '(')
		end := strings.IndexRune(s[1], ')')
		in := s[1][start+1:end]
		out := s[1][:start]
		if len(in) > 0 && in == strings.Repeat("9", len(in)) {
			countZero := preZero(out)
			outInt, _ := strconv.Atoi(out)
			outInt += 1
			newOut := strconv.Itoa(outInt)
			if len(newOut) > len(out) {
				if countZero > 0 {
					countZero--
				} else {
					x += 1
					newOut = "0"
				}
			}
			out = strings.Repeat("0", countZero) + newOut
			in = "0"
		}
			y = out + strings.Repeat(in, 100)
	}
	if len(y) >= 100 {
		y = y[:100]
	}
	return
}

func preZero(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			count++
		} else {
			break
		}
	}
	return count
}
