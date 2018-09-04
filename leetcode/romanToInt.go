package main

import "fmt"

func main() {
	ret := romanToInt("MCDXXXVIII")//CXCIX
	fmt.Println(ret)
}

func romanToInt(s string) int {
	map1 := map[string]int{}
	if len(s) == 0 {
		return 0
	}
	map1["I"] = 1
	map1["V"] = 5
	map1["X"] = 10
	map1["L"] = 50
	map1["C"] = 100
	map1["D"] = 500
	map1["M"] = 1000
	before := map1[string(s[0])]
	ret := before
	for i := 1; i < len(s); i++ {
		n := map1[string(s[i])]
		ret += n
		if before < n {
			ret -= before * 2
		}
		before = n
	}
	return ret
}
