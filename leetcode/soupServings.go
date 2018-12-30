package main

import (
	"fmt"
	"math"
)

func main() {
	n := 6000
	fmt.Println(soupServings1(n))
}

func soupServings(N int) float64 {
	var ret float64
	na := N
	for a := 0; ; a++ {
		nb := na - a * 100
		if nb <= 0 {
			fmt.Println("add a", a)
			ret += math.Pow(0.25, float64(a))
			break
		}
		for b := 0; ; b++ {
			nc := nb - b * 75
			if nc <= 0 {
				fmt.Println("add a b", a, b)
				ret += math.Pow(0.25, float64(a+b))
				break
			}
			for c := 0; ; c++ {
				nd := nc - c * 50
				if nd <= 0 {
					if a == 0 && b == 0 {
						fmt.Println("add a b c", a, b, c)
						ret += math.Pow(0.25, float64(a+b+c)) * 0.5
					} else {
						fmt.Println("add a b c 1", a, b, c)
						ret += math.Pow(0.25, float64(a+b+c))
					}
					break
				}
				for d := 0; ; d++ {
					ra := a * 100 + b * 75 + c * 50 + d * 25
					rb := b * 25 + c * 50 + d * 75
					if ra >= N && rb < N {
						fmt.Println("add a b c d", a, b, c, d)
						ret += math.Pow(0.25, float64(a+b+c+d))
					}
					if ra >= N || rb >= N {
						break
					}
				}
			}
		}
	}
	return ret
}

func soupServings1(N int) float64 {
	return ss(N, N, make(map[string]float64))
}

func ss(a, b int, m map[string]float64) float64 {
	if val, ok := m[fmt.Sprintf("%d+%d", a, b)]; ok {
		return val
	}
	if a <= 0 && b > 0 {
		return 1.0
	}
	if a <= 0 && b <= 0 {
		return 0.5
	}
	if a > 0 && b <= 0 {
		return 0
	}
	v := 0.25 * (ss(a-100, b, m) + ss(a-75, b-25, m) + ss(a-50, b-50, m) + ss(a-25, b-75, m))
	m[fmt.Sprintf("%d+%d", a, b)] = v
	return v
}
