package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcodle"))
}

func firstUniqChar(s string) int {

	lower := make([]int, 26)

	for i := 0; i < len(s); i++ {
		lower[s[i]-97] += 1
	}

	fmt.Println(lower)

	for i := 0; i < len(s); i++ {
		if lower[int(s[i]) - 97] == 1 {
			return i
		}
	}

	return -1
}
