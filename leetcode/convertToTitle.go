package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1231234))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(18278))
}

func convertToTitle(n int) string {
	ret := ""
	if n == 0 {
		return ""
	} else {
		ret = convertToTitle((n-1)/26) + string('A' + (n-1)%26)
		//fmt.Println(ret)
	}
	return ret
}
