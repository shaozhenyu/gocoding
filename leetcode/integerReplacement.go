package main

import "fmt"

func main() {
	fmt.Println(integerReplacement(35353))
}

func integerReplacement(n int) int {
	return intRe(n, make(map[int]int))
}

func intRe(n int, m map[int]int) (count int) {
	defer func(c int){
		m[c] = count
	}(n)

	if _, ok := m[n]; ok {
		return m[n]
	}
	for n%2 == 0 {
		n /= 2
		count++
		if _, ok := m[n]; ok {
			return m[n] + 1
		}
	}
	if n == 1 {
		return count
	}
	count1 := intRe(n+1, m)
	count2 := intRe(n-1, m)
	count += min(count1, count2) + 1
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
