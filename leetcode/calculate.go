package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calculate("(8+9)-(7+11) + (4 + 5 - (6 + 7) - (10 - 20))"))
}

func calculate(s string) int {
	sli := make([]string, len(s))
	pos := 0
	val := ""
	f := func() {
		if val != "" {
			sli[pos] = val
			pos++
			val = ""
		}
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '+', '-':
			f()
			sli[pos] = string(s[i])
			pos++
		case ')':
			f()
			var val string
			val, pos = cal(sli, pos)
			sli[pos] = val
			pos++
		case ' ':
			f()
		default:
			val += string(s[i])
		}
	}
	f()
	ret, _ := strconv.Atoi(do(sli[:pos]))
	return ret
}

func cal(s []string, pos int) (string, int) {
	for i := pos - 1; i >= 0; i-- {
		if s[i] == "(" {
			val := do(s[i+1:pos])
			p := i
			return val, p
		}
	}
	return "", 0
}

func do(s []string) string {
	ret := ""
	for i := 0; i < len(s); i++  {
		if s[i] == "+" {
			ret = add(ret, s[i+1])
			i++
		} else if s[i] == "-" {
			ret = sub(ret, s[i+1])
			i++
		} else {
			ret = s[i]
		}
	}
	return ret
}

func add(a, b string) string {
	i1, _ := strconv.Atoi(a)
	i2, _ := strconv.Atoi(b)
	return strconv.Itoa(i1+i2)
}


func sub(a, b string) string {
	i1, _ := strconv.Atoi(a)
	i2, _ := strconv.Atoi(b)
	return strconv.Itoa(i1-i2)
}

