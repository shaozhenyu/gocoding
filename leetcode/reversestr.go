package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	k, _ := strconv.Atoi(os.Args[2])
	fmt.Println(reverseStr(os.Args[1], k))
}

func reverseStr(s string, k int) string {
	if len(s) > 2 * k {
		return reverseStr(s[:2*k], k) + reverseStr(s[2*k:], k)
	}
	if len(s) > k {
		return reverse(s[:k]) + s[k:]
	}
	return reverse(s)
}

func reverse(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b)
}

func reverseStr1(s string, k int) string {
	b := []byte(s)
	lenS := len(s)
	if lenS <= k {
		for i := 0; i < lenS/2; i++ {
			b[i], b[lenS-i-1] = b[lenS-i-1], b[i]
		}
		return string(b)
	}

	start, end := 0, k-1
	fmt.Println(lenS)
	for start < lenS {
		if end >= lenS {
			end = lenS - 1
		}
		fmt.Println(start, end)
		for i := 0; i <= (end - start)/2; i++ {
			b[start+i], b[end-i] = b[end-i], b[start+i]
		}
		start += (k * 2)
		end += (k * 2)
	}
	return string(b)
}
