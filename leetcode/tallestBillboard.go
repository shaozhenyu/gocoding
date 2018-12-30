package main

import "fmt"

func main() {
	rods := []int{118,131,127,121,126,123,124,111,105,125,103,132,101,121,110,144,114,120,124}
	fmt.Println(tallestBillboard(rods))
}

func tallestBillboard(rods []int) int {
	sum := 0
	for i := 0; i < len(rods); i++ {
		sum += rods[i]
	}
	avg := sum/2
	m := make(map[int]map[int]bool)
	m[0] = make(map[int]bool)
	m[0][0] = true
	for i := 0; i < len(rods); i++ {
		newM := make(map[int]map[int]bool)
		for k, v := range m {
			for v1 := range v {
				if _, ok := newM[k]; !ok {
					newM[k] = make(map[int]bool)
				}
				newM[k][v1 << 1] = true
				if k + rods[i] > avg {
					continue
				}
				if _, ok := newM[k+ rods[i]]; !ok {
					newM[k + rods[i]] = make(map[int]bool)
				}
				newM[k + rods[i]][v1 << 1 + 1] = true
			}
		}
		m = newM
	}
	max := 0
	ccount := 0
	for k, v := range m {
		ccount += len(v)
		if len(v) < 2 || k <= max {
			continue
		}
		for v1 := range v {
			for v2 := range v {
				if v1 & v2 == 0 {
					max = k
					break
				}
			}
			if max == k {
				break
			}
		}
	}
	fmt.Println(m)
	fmt.Println(ccount)
	return max
}
