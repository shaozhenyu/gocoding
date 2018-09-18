package main

import "fmt"

func main() {
	str := "(00011)"
	fmt.Println(ambiguousCoordinates(str))
}

func ambiguousCoordinates(S string) []string {
	S = S[1 : len(S)-1]
	ret := make([]string, 0)
	for i := 0; i < len(S); i++ {
		ret1 := handle(S[:i+1], S[i+1:])
		ret = append(ret, ret1...)
	}
	return ret
}

func handle(s1, s2 string) []string {
	if s1 == "" || s2 == "" {
		return []string{}
	}
	fmt.Println("s1, s2: ", s1, s2)
	ret := make([]string, 0)
	p1 := split(s1)
	p2 := split(s2)
	if len(p1) == 0 || len(p2) == 0 {
		return ret
	}
	fmt.Println("p1, p2", p1, p2)
	for _, a := range p1 {
		for _, b := range p2 {
			ret = append(ret, fmt.Sprintf("(%s, %s)", a, b))
		}
	}
	return ret
}

// "00", "0.0", "0.00", "1.0", "001", "00.01"
// 点前面不能出现2个0， 点最后不能是0
func split(s string) []string {
	ret := []string{}
	if len(s) == 1 {
		ret = append(ret, s)
		return ret
	}
	for i := 1; i < len(s); i++ {
		pre := s[:i]
		if len(pre) >= 2 && pre[0] == '0' {
			continue
		}
		suf := s[i:]
		if suf[len(suf)-1] == '0' {
			continue
		}
		v := s[:i] + "." + s[i:]
		ret = append(ret, v)
	}
	if s[0] != '0' {
		ret = append(ret, s)
	}
	return ret
}
