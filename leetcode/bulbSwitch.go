package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(bulbSwitch(i))
	}
}

func bulbSwitch(n int) int {
	if n < 2 {
		return n
	}
	m := make(map[int]int)
	for i := 2; i < n; i++ {
		m[i] = 0
	}
	ret := 1
	for i := 2; i <= n; i++ {
		if m[i]%2 != 0 {
			ret++
		}
		times := 2
		for times * i <= n {
			m[times * i]++
			times++
		}
	}
	return ret
}
