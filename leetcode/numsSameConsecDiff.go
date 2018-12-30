package main

import (
	"fmt"
	// "math"
)

func main() {
	N := 9
	K := 1
	fmt.Println(numsSameConsecDiff(N, K))
}

func numsSameConsecDiff(N int, K int) []int {
	ret := []int{}
	ten := 1
	for i := 1; i < N; i++ {
		ten *= 10
	}
	for i := 1; i <= 9; i++ {
		v := next(N, K, i, ten)
		if len(v) > 0 {
			ret = append(ret, v...)
		}
	}
	return ret
}

func next(N, k int, now int, ten int) []int {
	if N == 1 {
		return []int{now}
	}
	ret := []int{}
	x := now * ten
	fmt.Println(x, k)
	a := now - k
	if a >= 0 {
		v := next(N-1, k, a, ten/10)
		for _, vv := range v {
			ret = append(ret, x + vv)
		}
	}
	b := now + k
	if b >= 0 && b < 10 {
		v := next(N-1, k, b, ten/10)
		for _, vv := range v {
			ret = append(ret, x + vv)
		}
	}
	return ret
}
