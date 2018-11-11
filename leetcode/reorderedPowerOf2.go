package main

import "fmt"

func main() {
	N := 100
	fmt.Println(reorderedPowerOf2(N))
}

var (
	s = initSource()
)

func initSource() map[int][]map[int]int {
	m := make(map[int][]map[int]int)
	start := 1
	for start <= 1000000000 {
		t := getM(start)
		size := len(t)
		if _, ok := m[size]; !ok {
			m[size] = make([]map[int]int, 0)
		}
		m[size] = append(m[size], t)
		start *= 2
	}
	return m
}

func reorderedPowerOf2(N int) bool {
	target := getM(N)
	if source, ok := s[len(target)]; ok {
		for i := 0; i < len(source); i++ {
			flag := true
			for k, v := range source[i] {
				if target[k] != v {
					flag = false
					break
				}
			}
			if flag {
				return true
			}
		}
	}
	return false
}

func getM(N int) map[int]int {
	target := make(map[int]int)
	for N > 0 {
		t := N%10
		target[t] += 1
		N /= 10
	}
	return target
}
