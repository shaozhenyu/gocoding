package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(10))
}

func str(n int) string {
	return strconv.Itoa(n)
}

func ff(n string) string {
	if len(n) == 1 {
		return ("1" + n)
	}

	s := ""
	last := string(n[0])
	length := 1

	for i := 1; i < len(n); i++ {
		if string(n[i]) == last {
			length++
		} else {
			s += str(length) + last
			length = 1
			last = string(n[i])
		}
	}

	s += str(length) + string(n[len(n)-1])

	return s
}

func countAndSay(n int) string {
	k := ""
	var m string
	for i := 1; i <= n; i++ {
		if i == 1 {
			k = "1"
		} else {
			m = ff(k)
			k = m
			fmt.Println(i, k)
		}
	}

	return k
}
