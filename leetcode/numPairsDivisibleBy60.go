package main

import "fmt"

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{439,407,197,191,291,486,30,307,11}))
}

func numPairsDivisibleBy60(time []int) int {
	m := make(map[int]int)
	for i := 0; i < len(time); i++ {
		m[time[i]%60]++
	}
	sum := 0
	fmt.Println(m)
	if m[0] > 1 {
		sum += m[0] * (m[0] - 1)/2
	}
	delete(m, 0)
	if m[30] > 1 {
		sum += m[30] * (m[30] - 1)/2
	}
	delete(m, 30)
	for k, v := range m {
		if v1, ok := m[60-k]; ok {
			sum += v * v1
			delete(m, 60-k)
		}
		delete(m, k)
	}
	return sum
}
