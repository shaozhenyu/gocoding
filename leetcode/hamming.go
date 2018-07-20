package main

import "fmt"

func main() {
	x := 123345
	y := 12331
	fmt.Println(hammingDistance(x, y))
}

func hammingDistance(x int, y int) int {
	ret := 0
	X := fmt.Sprintf("%b", x)
	Y := fmt.Sprintf("%b", y)
	if len(X) > len(Y) {
		X, Y = Y, X
	}
	preY := Y[:len(Y)-len(X)]
	sufY := Y[len(Y)-len(X):]
	for i := 0; i < len(X); i++{
		if X[i] != sufY[i] {
			ret++
		}
	}
	for _, p := range preY {
		if p == '1' {
			ret++
		}
	}
	return ret
}

