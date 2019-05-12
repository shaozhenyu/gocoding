package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(queryString("0110", 1))
	fmt.Println(queryString("0110", 2))
	fmt.Println(queryString("0110", 3))
	fmt.Println(queryString("0110", 4))
	fmt.Println(queryString("0110", 5))
	fmt.Println(queryString("0110", 14))
	fmt.Println(queryString("0110", 45))
}

func queryString(S string, N int) bool {
	k := 1
	for k < N {
		k *= 2
	}
	if k == N {
		k /= 2
	} else {
		k /= 4
	}
	if k == 0 {
		k = 1
	}
	fmt.Println(k, N)
	for i := N; i >= k; i-- {
		str := fmt.Sprintf("%b", i)
		if !strings.Contains(S, str) {
			return false
		}
	}
	return true
}
