package main

import "fmt"

func main() {
	fmt.Println(divisorGame(3))
}

func divisorGame(N int) bool {
	a, b := make(map[int]bool), make(map[int]bool)
	return dg(N, a, b)
}

func dg(N int, a, b map[int]bool) bool {
	if _, ok := a[N]; ok {
		return a[N]
	}
	for i := 1; i < N; i++ {
		if N%i == 0 {
			if !dg(N-i, b, a) {
				a[N] = true
				return true
			}
		}
	}
	a[N] = false
	return false
}
