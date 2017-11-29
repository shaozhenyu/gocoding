package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(evaluate("(let x 2 y (mult x (add x 5)) (add 2 y))"))
	//fmt.Println(evaluate("(add (mult 2 (add 2 5)) 5)"))
	//fmt.Println(evaluate("(mult 5 6)"))
	fmt.Println(evaluate("(let var 78 b 77 (let c 33 (add c (mult var 66))))"))
}

func evaluate(expression string) int {
	m := make(map[string]int)
	return eval(expression, m)
}

func eval(expression string, m map[string]int) int {
	s := parse(expression[1 : len(expression)-1])

	if s[0] == "add" {
		return cal(s[1], m) + cal(s[2], m)
	}

	if s[0] == "mult" {
		return cal(s[1], m) * cal(s[2], m)
	}

	if s[0] == "let" {
		newM := make(map[string]int)
		for key, v := range m {
			newM[key] = v
		}
		for i := 1; i < len(s)-1; i += 2 {
			newM[s[i]] = cal(s[i+1], newM)
		}
		return cal(s[len(s)-1], newM)
	}

	return 0
}

func cal(v string, m map[string]int) int {
	fmt.Println("cal:", v, m)
	var ret int
	if v[0] == byte('(') {
		return eval(v, m)
	} else {
		if val, ok := m[v]; ok {
			ret = val
		} else {
			ret, _ = strconv.Atoi(v)
		}
	}
	return ret
}

func parse(str string) []string {
	start := 0
	pair := 0
	s := []string{}
	for i := 1; i < len(str)-1; i++ {
		if pair == 0 && str[i] == byte(' ') {
			s = append(s, str[start:i])
			start = i + 1
		} else if str[i] == byte('(') {
			pair++
		} else if str[i] == byte(')') {
			pair--
		}
	}

	s = append(s, str[start:])

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println("----------------------------")
	return s
}
