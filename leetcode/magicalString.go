package main

import "fmt"

func main() {
	fmt.Println(magicalString(10))
}

func magicalString(n int) int {
	str := "122"
	if n == 0 {
		return 0
	}
	count := 1
	now := "1"
	idx := 2
	end := 2
	for len(str) < n {
		if str[idx] == '1' {
			str += now
		} else {
			str += now + now
		}
		if end < len(str) && str[end+1] == '1' {
			end++
			count++
		}
		idx++
		if now == "1" {
			now = "2"
		} else {
			now = "1"
		}
	}
	return count
}
