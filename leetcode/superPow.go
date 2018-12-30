package main

import "fmt"

func main() {
	a := 2
	b := []int{1,3,4,3,2}
	fmt.Println(superPow(a, b))
}

func superPow(a int, b []int) int {
	if a == 1 {
		return 1
	}
	m := make(map[int]bool)
	s := make([]int, 1338)
	val := a
	for i := 0; i <= 1338; i++ {
		val = val%1337
		if m[val] {
			break
		}
		m[val] = true
		s[i] = val
		val = a * val
	}
	v := 0
	for i := 0; i < len(b); i++ {
		v = (v * 10 + b[i])%len(m)
	}
	return s[v%len(m)-1]
}
