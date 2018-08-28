package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
}

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, c := range cpdomains {
		s := strings.Split(c, " ")
		num, _ := strconv.Atoi(s[0])
		childs := strings.Split(s[1], ".")
		name := ""
		for i := len(childs)-1; i >= 0; i-- {
			name = childs[i] + "." + name
			m[name] += num
		}
	}
	ret := make([]string, 0, len(m))
	for k, v := range m {
		ret = append(ret, fmt.Sprintf("%d %s", v, k[:len(k)-1]))
	}
	return ret
}
