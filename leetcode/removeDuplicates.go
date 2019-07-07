package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {
	fmt.Println(S)
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			fmt.Println("xxx")
			if i + 1 < len(S) {
				if S[i+1] == S[i] {
					return removeDuplicates(S[:i-1] + S[i+2:])
				}
			}
			return removeDuplicates(S[:i-1] + S[i+1:])
		}
	}
	return S
}
