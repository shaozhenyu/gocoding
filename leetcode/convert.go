package main

import "fmt"

func main() {
	s := convert("AB", 1)
	fmt.Println(s)
}

func convert(s string, numRows int) string {

	arr := make([]string, len(s))

	j := 0
	for i := 0; i < numRows; i++ {
		k := i
		for ; j < len(s); j++ {
			if k >= len(s) {
				break
			}
			if i == 0 || i == numRows-1 {
				arr[j] = string(s[k])
			} else {
				arr[j] = string(s[k])
				j++
				mid := k + (numRows - (i + 1)) * 2
				if mid >= len(s) {
					break
				}
				arr[j] = string(s[mid])
			}
			k = k + numRows * 2 - 2
		}
		fmt.Println(arr)
	}

	ret := ""
	for _, v := range arr {
		ret += v
	}
	return ret
}
