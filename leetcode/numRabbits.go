package main

import "fmt"

func main() {
	a := []int{1,1,2}
	fmt.Println(numRabbits(a))
}

func numRabbits(answers []int) int {
	m := make(map[int]int)
	for _, a := range answers {
		m[a]++
	}
	count := 0
	for k, v := range m {
		for v > 0 {
			count += k + 1
			v = v - k - 1
		}
	}
	return count
}
