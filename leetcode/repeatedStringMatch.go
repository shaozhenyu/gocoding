package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "c"
	B := "c"
	A = "abababaaba"
	B = "aabaaba"
	fmt.Println(repeatedStringMatch(A, B))
}

func repeatedStringMatch(A string, B string) int {
	if strings.Contains(A, B) {
		return 1
	}
	for i := 0; i < len(A); i++ {
		if A[i] == B[0] {
			if len(B) < len(A[i:]) {
				continue
			}
			fmt.Println(A[i:], B[:len(A[i:])])
			if A[i:] != B[:len(A[i:])] {
				continue
			}
			ret := match(A, B[len(A[i:]):])
			if ret >= 0 {
				return ret + 1
			}
		}
	}
	return -1
}

func match(A string, B string) int {
	if len(B) == 0 {
		return 0
	}
	count := len(B)/len(A)
	if len(B)%len(A) > 0 {
		count += 1
	}
	fmt.Println(A, B)
	A = strings.Repeat(A, count)
	if B == A[:len(B)] {
		return count
	}
	return -1
}
