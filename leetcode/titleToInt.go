package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("AB"))
}

func titleToNumber(s string) int {
	n := 1
	sum := 0
	for i := len(s)-1; i >= 0; i-- {
		sum += (int(s[i]) - 64) * n
		n = n * 26
	}
	return sum
}
